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

func (i *Indexer) Run() error {
	client, err := ethclient.Dial(i.cfg.InfuraURL)
	if err != nil {
		i.logger.Error().Err(err).Msgf("Failed to dial the Ethereum client: %v", err)
		return err
	}

	i.logger.Debug().Msgf("Connected to Ethereum client: %s", i.cfg.InfuraURL)

	chainID, err := client.ChainID(i.ctx)
	if err != nil {
		i.logger.Error().Err(err).Msg("Failed to fetch chain ID")
		return err
	}

	i.logger.Info().Msgf("Connected to Ethereum chain ID: %s", chainID.String())

	subscriber := NewSubscriber(
		i.logger,
		i.ctx,
		i.cfg,
		client,
		i.blockRepo,
		i.transactionRepo,
		i.eventLogRepo,
		i.checkpointRepo,
	)
	subscriber.SubscribeNewBlocks()

	return nil
}
