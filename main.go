package main

import (
	"context"
	"sync"

	"github.com/conceptcodes/eth-indexer-go/cmd/api"
	"github.com/conceptcodes/eth-indexer-go/cmd/indexer"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"

	"github.com/conceptcodes/eth-indexer-go/config"
	"github.com/conceptcodes/eth-indexer-go/pkg/email"
	"github.com/conceptcodes/eth-indexer-go/pkg/logger"
	"github.com/conceptcodes/eth-indexer-go/pkg/storage/db"
	"github.com/conceptcodes/eth-indexer-go/pkg/storage/redis"
)

func main() {
	cfg := config.LoadConfig()
	ctx := context.Background()

	log := logger.New()
	db := db.NewDB(cfg, log)
	redis := redis.NewRedisClient(cfg, log, ctx)
	emailClient := email.NewEmailClient(cfg, log)

	var wg sync.WaitGroup
	wg.Add(2)

	log.Info().Msg("Connecting to the database...")
	db.ConnectDB()

	log.Info().Msg("Connecting to redis...")
	redis.Connect()

	dbClient := db.GetDB()

	blockRepo := repository.NewBlockRepository(dbClient)
	transactionRepo := repository.NewTransactionRepository(dbClient)
	userRepo := repository.NewUserRepository(dbClient)
	authRepo := repository.NewAuthRepository(dbClient)
	eventLogRepo := repository.NewEventLogRepository(dbClient)
	checkpointRepo := repository.NewCheckPointRepository(dbClient)

	indexer := indexer.NewIndexer(cfg, ctx, log, blockRepo, transactionRepo, eventLogRepo, checkpointRepo)
	api := api.NewApi(
		cfg,
		log,
		db,
		transactionRepo,
		userRepo,
		authRepo,
		blockRepo,
		eventLogRepo,
		ctx,
		redis,
		emailClient,
	)

	go func() {
		defer wg.Done()
		log.Debug().Msg("Starting Indexer")
		if err := indexer.Run(); err != nil {
			log.Fatal().Err(err).Msgf("Indexer failed: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		log.Debug().Msg("Starting API")
		if err := api.Start(); err != nil {
			log.Fatal().Err(err).Msgf("API server failed: %v", err)
		}
	}()

	wg.Wait()
}
