package indexer

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/conceptcodes/eth-indexer-go/config"
	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"

	"github.com/rs/zerolog"
)

var blockQueue = make(chan uint64, 100)

type Subscriber struct {
	log             *zerolog.Logger
	cfg             *config.Config
	ctx             context.Context
	client          *ethclient.Client
	blockRepo       repository.BlockRepository
	transactionRepo repository.TransactionRepository
	eventLogRepo    repository.EventLogRepository
	checkpointRepo  repository.CheckpointRepository
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
	if err != nil || checkpoint == nil {
		s.log.Warn().Msg("No checkpoint found, starting from configured block number")
		return s.cfg.BlockNumber
	}
	return checkpoint.LastBlock
}

func (s *Subscriber) processHistoricalBlocks(fromBlock, toBlock uint64) {
	blockQueue := make(chan uint64, 10)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
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

	for i := 0; i < 5; i++ {
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
			return
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
	maxRetries := 3
	retryDelay := time.Second

	for i := 0; i < maxRetries; i++ {
		block, err = s.client.BlockByNumber(s.ctx, big.NewInt(int64(blockNumber)))
		if err == nil && block != nil {
			break
		}
		s.log.Warn().Err(err).Msgf("Retrying block %d fetch (%d/%d)", blockNumber, i+1, maxRetries)
		time.Sleep(retryDelay)
		retryDelay *= 2 
	}

	if err != nil || block == nil {
		s.log.Error().Err(err).Msgf("Failed to fetch block %d after retries", blockNumber)
		return
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
	}

	if err := s.blockRepo.Create(&blockData); err != nil {
		s.log.Error().Err(err).Msg("Failed to save block to repository")
	}
	s.log.Debug().Msgf("Stored block %d", block.NumberU64())
}

func (s *Subscriber) processTransactions(block *types.Block) {
	var transactions []*models.Transaction

	chainID, err := s.client.NetworkID(s.ctx)
	if err != nil {
		s.log.Fatal().Err(err).Msg("Failed to get network Chain ID")
	}

	for _, tx := range block.Transactions() {
		if tx == nil {
			s.log.Error().Msg("Transaction is nil, skipping")
			continue
		}

		if tx.ChainId() == nil {
			s.log.Warn().Msgf("Skipping unsigned transaction: %s", tx.Hash().Hex())
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
			s.log.Warn().Err(err).Msgf("Failed to derive sender for tx: %s", tx.Hash().Hex())
			continue
		}

		toAddress := "Contract Creation"
		if tx.To() != nil {
			toAddress = tx.To().Hex()
		}

		transaction := models.Transaction{
			Hash:        tx.Hash().Hex(),
			BlockNumber: block.NumberU64(),
			From:        msg.Hex(),
			To:          toAddress,
			Value:       tx.Value().String(),
			GasPrice:    tx.GasPrice().String(),
			GasLimit:    tx.Gas(),
		}

		transactions = append(transactions, &transaction)
	}

	s.transactionRepo.CreateAll(transactions)
	s.log.Debug().Msgf("Processed %d transactions in block %d", len(transactions), block.NumberU64())
}

func (s *Subscriber) processEvents(block *types.Block) {
	blockHash := block.Hash()
	query := ethereum.FilterQuery{
		BlockHash: &blockHash,
	}

	logs, err := s.client.FilterLogs(context.Background(), query)
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to fetch logs")
		return
	}

	events := make([]*models.Event, len(logs))
	for i, vLog := range logs {
		events[i] = &models.Event{
			LogIndex:        vLog.Index,
			TransactionHash: vLog.TxHash.Hex(),
			BlockNumber:     block.NumberU64(),
			Address:         vLog.Address.Hex(),
			Data:            sanitizeData(vLog.Data),
		}
	}

	s.eventLogRepo.CreateAll(events)
	s.log.Debug().Msgf("Processed %d events", len(events))
}

func (s *Subscriber) updateCheckpoint(blockNumber uint64) {
	checkpoint := models.Checkpoint{LastBlock: blockNumber}
	err := s.checkpointRepo.Create(&checkpoint)
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to update checkpoint")
	}
	s.log.Debug().Msgf("Checkpoint updated to block %d", blockNumber)
}

func sanitizeData(input []byte) string {
	str := strings.TrimPrefix(string(input), "0x")
	return strings.TrimSpace(str)
}
