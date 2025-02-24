package indexer

import (
	"context"

	"github.com/conceptcodes/eth-indexer-go/config"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"
	"github.com/rs/zerolog"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Indexer struct {
	cfg             *config.Config
	ctx             context.Context
	logger          *zerolog.Logger
	blockRepo       repository.BlockRepository
	transactionRepo repository.TransactionRepository
	eventLogRepo    repository.EventLogRepository
	checkpointRepo  repository.CheckpointRepository
	client          *ethclient.Client
}

func NewIndexer(
	cfg *config.Config,
	ctx context.Context,
	logger *zerolog.Logger,
	blockRepo repository.BlockRepository,
	transactionRepo repository.TransactionRepository,
	eventLogRepo repository.EventLogRepository,
	checkpointRepo repository.CheckpointRepository,
) *Indexer {
	return &Indexer{
		cfg:             cfg,
		ctx:             ctx,
		logger:          logger,
		blockRepo:       blockRepo,
		transactionRepo: transactionRepo,
		eventLogRepo:    eventLogRepo,
		checkpointRepo:  checkpointRepo,
	}
}

func (i *Indexer) connect() error {
	var lastErr error
	for _, url := range i.cfg.RpcUrls {
		c, err := ethclient.Dial(url)
		if err != nil {
			i.logger.Error().Err(err).Msgf("Failed to dial Ethereum client at %s", url)
			lastErr = err
			continue
		}

		chainID, err := c.ChainID(i.ctx)
		if err != nil {
			i.logger.Error().Err(err).Msgf("Failed to fetch chain ID from %s", url)
			lastErr = err
			continue
		}

		i.logger.Info().Msgf("Connected to Ethereum chain ID: %s using RPC Url: %s", chainID.String(), url)
		i.client = c
		return nil
	}

	return lastErr
}

func (i *Indexer) Run() error {
	if err := i.connect(); err != nil {
		i.logger.Fatal().Err(err).Msg("All RPC URLs exhausted")
		return err
	}

	subscriber := NewSubscriber(
		i.logger,
		i.ctx,
		i.cfg,
		i.client,
		i.blockRepo,
		i.transactionRepo,
		i.eventLogRepo,
		i.checkpointRepo,
		i.connect,
	)
	subscriber.StartIndexing()

	return nil
}
