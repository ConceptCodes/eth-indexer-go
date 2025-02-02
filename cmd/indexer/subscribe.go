package indexer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/conceptcodes/eth-indexer-go/config"
	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"

	"github.com/rs/zerolog"
)

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
	for blockNumber := fromBlock; blockNumber <= toBlock; blockNumber++ {
		s.processBlock(blockNumber)
	}
}

func (s *Subscriber) SubscribeNewBlocks() {
	headers := make(chan *types.Header)
	sub, err := s.client.SubscribeNewHead(s.ctx, headers)
	if err != nil || sub == nil {
		s.log.Fatal().Err(err).Msg("Failed to subscribe to new blocks")
	}
	s.log.Info().Msg("Subscribed to new block headers...")

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
			go s.processBlock(header.Number.Uint64())
		}
	}
}

func (s *Subscriber) processBlock(blockNumber uint64) {
	block, err := s.client.BlockByNumber(s.ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		s.log.Error().Err(err).Msgf("Failed to fetch block %d", blockNumber)
		return
	}

	if block == nil {
		s.log.Error().Msgf("Block %d is nil", blockNumber)
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

	s.log.Info().Str("block", block.Hash().Hex()).Msgf("Processing block: %s", block.Hash().Hex())
	s.log.Debug().Msgf("Block has %d transactions", len(block.Transactions()))

	for _, tx := range block.Transactions() {
		if tx == nil {
			s.log.Error().Msg("Transaction is nil")
			continue
		}
		msg, err := types.NewEIP155Signer(tx.ChainId()).Sender(tx)

		if err != nil {
			s.log.Error().Err(err).Msgf("Failed to derive sender from transaction: %s", tx.Hash().Hex())
		}

		s.log.Info().Msgf("Sender: %s", msg.Hex())

		toAddress := "Contract Creation"
		if tx.To() != nil {
			toAddress = tx.To().Hex()
		}

		s.log.
			Debug().
			Str("hash", tx.Hash().Hex()).
			Str("from", msg.Hex()).
			Str("block_number", block.Number().String()).
			Str("to", toAddress).
			Str("value", tx.Value().String()).
			Str("gas_price", tx.GasPrice().String()).
			Uint64("gas_limit", tx.Gas()).
			Msgf("Processing transaction: %s", tx.Hash().Hex())

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
	s.log.Debug().Msgf("Processed %d transactions", len(transactions))
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
			Data:            string(vLog.Data),
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
