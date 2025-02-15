package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/conceptcodes/eth-indexer-go/config"
	"github.com/conceptcodes/eth-indexer-go/internal/constants"
	"github.com/conceptcodes/eth-indexer-go/internal/handlers"
	"github.com/conceptcodes/eth-indexer-go/internal/helpers"
	"github.com/conceptcodes/eth-indexer-go/internal/middleware"
	"github.com/conceptcodes/eth-indexer-go/internal/repository"
	"github.com/conceptcodes/eth-indexer-go/pkg/email"
	"github.com/conceptcodes/eth-indexer-go/pkg/storage/db"
	"github.com/conceptcodes/eth-indexer-go/pkg/storage/redis"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Api struct {
	cfg             *config.Config
	logger          *zerolog.Logger
	db              *db.DBConfig
	userRepo        repository.UserRepository
	authRepo        repository.AuthRepository
	blockRepo       repository.BlockRepository
	transactionRepo repository.TransactionRepository
	eventLogRepo    repository.EventLogRepository
	ctx             context.Context
	redisClient     *redis.RedisClient
	emailClient     *email.EmailClient
}

func NewApi(
	cfg *config.Config,
	logger *zerolog.Logger,
	db *db.DBConfig,
	transactionRepo repository.TransactionRepository,
	userRepo repository.UserRepository,
	authRepo repository.AuthRepository,
	blockRepo repository.BlockRepository,
	eventLogRepo repository.EventLogRepository,
	ctx context.Context,
	redisClient *redis.RedisClient,
	emailClient *email.EmailClient,
) *Api {
	return &Api{
		cfg:             cfg,
		logger:          logger,
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
		authRepo:        authRepo,
		eventLogRepo:    eventLogRepo,
		blockRepo:       blockRepo,
		ctx:             ctx,
		redisClient:     redisClient,
		emailClient:     emailClient,
	}
}

func (a *Api) Start() error {
	router := mux.NewRouter()

	authHelper := helpers.NewAuthHelper(a.logger, a.userRepo)
	responseHelper := helpers.NewResponseHelper(a.logger)
	validatorHelper := helpers.NewValidatorHelper(a.logger, responseHelper)

	redisClient := a.redisClient.GetClient()
	redisHelper := helpers.NewRedisHelper(redisClient, a.logger, a.ctx)

	router.Use(middleware.ContentTypeJSON)

	traceMiddleware := middleware.NewTraceRequestMiddleware(a.logger, authHelper, responseHelper)
	router.Use(traceMiddleware.Start)

	requestLogger := middleware.NewLoggerMiddleware(a.logger)
	router.Use(requestLogger.Start)

	rateLimitMiddleware := middleware.NewRateLimitRequestMiddleware(a.logger, redisClient, a.cfg)
	router.Use(rateLimitMiddleware.Start)

	userHandler := handlers.NewUserHandler(
		a.userRepo,
		a.authRepo,
		a.logger,
		authHelper,
		responseHelper,
		validatorHelper,
		a.emailClient,
		redisHelper,
	)

	healthHandler := handlers.NewHealthHandler(a.logger, a.db, a.redisClient, responseHelper)
	blockHandler := handlers.NewBlockHandler(a.logger, a.blockRepo, responseHelper)
	transactionHandler := handlers.NewTransactionHandler(a.transactionRepo, a.logger, responseHelper)
	eventsHandler := handlers.NewEventHandler(a.logger, a.eventLogRepo, responseHelper)

	router.HandleFunc(constants.RegisterEndpoint, userHandler.RegisterUserHandler).Methods(http.MethodPost)
	router.HandleFunc(constants.LoginEndpoint, userHandler.LoginUserHandler).Methods(http.MethodPost)
	router.HandleFunc(constants.ForgotPasswordEndpoint, userHandler.ForgotPasswordHandler).Methods(http.MethodPost)
	router.HandleFunc(constants.ResetPasswordEndpoint, userHandler.ResetPasswordHandler).Methods(http.MethodPost)

	router.HandleFunc(constants.HealthCheckEndpoint, healthHandler.GetHealthHandler).Methods(http.MethodGet)
	router.HandleFunc(constants.ReadinessEndpoint, healthHandler.GetStatusHandler).Methods(http.MethodGet)

	router.HandleFunc(constants.GetBlockByNumberEndpoint, blockHandler.GetBlockByNumberHandler).Methods(http.MethodGet)
	router.HandleFunc(constants.GetTransactionsByHashEndpoint, transactionHandler.GetTransactionByHashHandler).Methods(http.MethodGet)
	router.HandleFunc(constants.GetEventsByContractAddressEndpoint, eventsHandler.GetEventLogsByAddressHandler).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", a.cfg.Host, a.cfg.Port),
		WriteTimeout: time.Duration(a.cfg.Timeout) * time.Second,
		ReadTimeout:  time.Duration(a.cfg.Timeout) * time.Second,
	}

	a.logger.Info().Msgf("Starting API server at %s:%s", a.cfg.Host, a.cfg.Port)
	return srv.ListenAndServe()
}
