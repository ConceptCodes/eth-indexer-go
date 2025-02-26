package indexer

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
	"gorm.io/gorm"
	"github.com/rs/zerolog"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/conceptcodes/eth-indexer-go/config"
	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"

)

var blockQueue = make(chan uint64, 100)
var maxRetries = 5

type Subscriber struct {
	log             *zerolog.Logger
	cfg             *config.Config
	ctx             context.Context
	client          *ethclient.Client
	blockRepo       repository.BlockRepository
	transactionRepo repository.TransactionRepository
	eventLogRepo    repository.EventLogRepository
	checkpointRepo  repository.CheckpointRepository
	connect         func() error
}

func NewSubscriber(
	log *zerolog.Logger,
	ctx context.Context,
	cfg *config.Config,
	client *ethclient.Client,
	blockRepo repository.BlockRepository,
	transactionRepo repository.TransactionRepository,
	eventLogRepo repository.EventLogRepository,
	checkpointRepo repository.CheckpointRepository,
	connect func() error,
) *Subscriber {
	return &Subscriber{
		log:             log,
		ctx:             ctx,
		cfg:             cfg,
		client:          client,
		blockRepo:       blockRepo,
		transactionRepo: transactionRepo,
		eventLogRepo:    eventLogRepo,
		checkpointRepo:  checkpointRepo,
		connect:         connect,
	}
}

func (s *Subscriber) StartIndexing() {
	lastIndexedBlock := s.getLastIndexedBlock()
	s.log.Info().Uint64("start_block", lastIndexedBlock).Msg("Starting indexing from block")

	header, err := s.client.HeaderByNumber(s.ctx, nil)
	if err != nil {
		s.log.Fatal().Err(err).Msg("Failed to fetch latest block header")
	}
	latestBlock := header.Number.Uint64()

	if lastIndexedBlock < latestBlock {
		s.processHistoricalBlocks(lastIndexedBlock, latestBlock)
	}

	s.SubscribeNewBlocks()
}

func (s *Subscriber) getLastIndexedBlock() uint64 {
	checkpoint, err := s.checkpointRepo.FindLastBlock()
	if err != nil || checkpoint == nil || checkpoint.LastBlock == 0 {
		s.log.Debug().Msg("No checkpoint found, starting from configured block number")
		return s.cfg.BlockNumber
	}
	return checkpoint.LastBlock
}

func (s *Subscriber) processHistoricalBlocks(fromBlock, toBlock uint64) {
	blockQueue := make(chan uint64, 10)
	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for blockNumber := range blockQueue {
				s.processBlock(blockNumber)
			}
		}()
	}

	for blockNumber := fromBlock; blockNumber <= toBlock; blockNumber++ {
		blockQueue <- blockNumber
	}
	close(blockQueue)

	wg.Wait()
}

func (s *Subscriber) SubscribeNewBlocks() {
	headers := make(chan *types.Header)
	sub, err := s.client.SubscribeNewHead(s.ctx, headers)
	if err != nil || sub == nil {
		s.log.Fatal().Err(err).Msg("Failed to subscribe to new blocks")
	}
	s.log.Info().Msg("Subscribed to new block headers...")

	for range 5 {
		go func() {
			for blockNumber := range blockQueue {
				s.processBlock(blockNumber)
			}
		}()
	}

	for {
		select {
		case err := <-sub.Err():
			s.log.Error().Err(err).Msg("Subscription error")
			if err := s.connect(); err != nil {
				s.log.Fatal().Err(err).Msg("Failed to reconnect to Ethereum client")
			}
			continue
		case header := <-headers:
			if header == nil {
				s.log.Error().Msg("Received nil header")
				continue
			}
			blockQueue <- header.Number.Uint64()
		}
	}
}

func (s *Subscriber) processBlock(blockNumber uint64) {
	var block *types.Block
	var err error
	retryDelay := time.Second

	// add a delay to avoid overwhelming the node
	time.Sleep(3 * time.Second)

	for i := range maxRetries {
		block, err = s.client.BlockByNumber(s.ctx, big.NewInt(int64(blockNumber)))
		if err == nil && block != nil {
			break
		}
		s.log.Warn().Err(err).Msgf("Retrying block %d fetch (%d/%d)", blockNumber, i+1, maxRetries)
		time.Sleep(retryDelay)
		retryDelay *= 2
	}

	if err != nil || block == nil {
		s.log.Fatal().Err(err).Msgf("Failed to fetch block %d after retries", blockNumber)
	}

	s.storeBlock(block)
	s.processTransactions(block)
	s.processEvents(block)
	s.updateCheckpoint(blockNumber)
}

func (s *Subscriber) storeBlock(block *types.Block) {
	blockData := models.Block{
		Number:     block.NumberU64(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		Size:       block.Size(),
		Miner:      block.Coinbase().Hex(),
		Timestamp:  block.Time(),
		TxHash:     block.TxHash().Hex(),
		Difficulty: block.Difficulty().Int64(),
		GasUsed:    block.GasUsed(),
		GasLimit:   block.GasLimit(),
		BaseFee:    block.BaseFee().String(),
	}

	if err := s.blockRepo.Create(&blockData); err != nil {
		s.log.Error().Err(err).Msg("Failed to save block to repository")
	}
	s.log.Info().Msgf("Stored block %d", block.NumberU64())
}

func (s *Subscriber) processTransactions(block *types.Block) {
	var transactions []*models.Transaction

	chainID, err := s.client.NetworkID(s.ctx)
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to get network Chain ID")
		if err := s.connect(); err != nil {
			s.log.Fatal().Err(err).Msg("Failed to reconnect to Ethereum client")
		}
		return
	}

	for _, tx := range block.Transactions() {
		if tx == nil {
			s.log.Error().Msg("Transaction is nil, skipping")
			continue
		}

		if tx.ChainId() == nil {
			s.log.Debug().Msgf("Skipping unsigned transaction: %s", tx.Hash().Hex())
			continue
		}

		var signer types.Signer
		if tx.Type() == types.DynamicFeeTxType {
			signer = types.NewLondonSigner(chainID) // EIP-1559
		} else {
			signer = types.NewEIP155Signer(chainID) // Legacy
		}

		msg, err := signer.Sender(tx)
		if err != nil {
			s.log.Debug().Err(err).Msgf("Failed to derive sender for tx: %s", tx.Hash().Hex())
			continue
		}

		toAddress := "Contract Creation"
		if tx.To() != nil {
			toAddress = tx.To().Hex()
		}

		receipt, err := s.client.TransactionReceipt(s.ctx, tx.Hash())
		if err != nil {
			s.log.Warn().Err(err).Msgf("Failed to fetch receipt for tx: %s", tx.Hash().Hex())
			continue
		}

		transaction := models.Transaction{
			Hash:        tx.Hash().Hex(),
			BlockNumber: block.NumberU64(),
			From:        msg.Hex(),
			To:          toAddress,
			Value:       tx.Value().String(),
			GasPrice:    tx.GasPrice().String(),
			GasLimit:    tx.Gas(),
			GasUsed:     receipt.GasUsed,
			Nonce:       tx.Nonce(),
			Timestamp:   block.Time(),
			Success:     types.ReceiptStatusSuccessful == receipt.Status,
		}


		transactions = append(transactions, &transaction)
	}
	err = s.transactionRepo.CreateAll(transactions)

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		s.log.Debug().Msgf("Skipping duplicate transaction(s) in block %d", block.NumberU64())
		return
	} else if err != nil {
		s.log.Fatal().Err(err).Msg("Failed to save transactions to repository")
		return
	}
	s.log.Info().Msgf("Processed %d transaction(s) in block %d", len(transactions), block.NumberU64())
}

func (s *Subscriber) processEvents(block *types.Block) {
	blockHash := block.Hash()
	query := ethereum.FilterQuery{
		BlockHash: &blockHash,
	}

	logs, err := s.client.FilterLogs(context.Background(), query)
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to fetch logs")
		if err := s.connect(); err != nil {
			s.log.Fatal().Err(err).Msg("Failed to reconnect to Ethereum client")
		}
		return
	}

	events := make([]*models.Event, len(logs))
	for i, vLog := range logs {
		events[i] = &models.Event{
			LogIndex:        vLog.Index,
			TransactionHash: vLog.TxHash.Hex(),
			BlockNumber:     block.NumberU64(),
			Address:         vLog.Address.Hex(),
			Topics:          mapTopics(vLog.Topics),
			// Data:            sanitizeData(vLog.Data),
		}
	}

	if err = s.eventLogRepo.CreateAll(events); err != nil {
		s.log.Error().Err(err).Msg("Failed to save events to repository")
		return
	}
	s.log.Debug().Msgf("Processed %d event(s)", len(events))
}

func mapTopics(topics []common.Hash) []string {
	mappedTopics := make([]string, len(topics))
	for i, topic := range topics {
		mappedTopics[i] = topic.Hex()
	}
	return mappedTopics
}

func (s *Subscriber) updateCheckpoint(blockNumber uint64) {
	checkpoint := models.Checkpoint{LastBlock: blockNumber}
	if err := s.checkpointRepo.Create(&checkpoint); err != nil {
		s.log.Error().Err(err).Msg("Failed to update checkpoint")
	}
	s.log.Debug().Msgf("Checkpoint updated to block %d", blockNumber)
}

func sanitizeData(input []byte) string {
	str := string(input)
	str = strings.ReplaceAll(str, "\x00", "")

	str = strings.TrimPrefix(str, "0x")
	str = strings.TrimSpace(str)

	if !utf8.ValidString(str) {
		validStr := make([]rune, 0, len(str))
		for i, r := range str {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(str[i:])
				if size == 1 {
					continue
				}
			}
			validStr = append(validStr, r)
		}
		str = string(validStr)
	}

	return str
}
