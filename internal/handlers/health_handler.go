package handlers

import (
	"net/http"

	"github.com/conceptcodes/eth-indexer-go/internal/helpers"
	"github.com/conceptcodes/eth-indexer-go/pkg/storage/db"
	"github.com/conceptcodes/eth-indexer-go/pkg/storage/redis"

	"github.com/rs/zerolog"
)

type HealthHandler struct {
	log            *zerolog.Logger
	dbClient       *db.DBConfig
	redisClient    *redis.RedisClient
	responseHelper *helpers.ResponseHelper
}

func NewHealthHandler(
	log *zerolog.Logger,
	db *db.DBConfig,
	redis *redis.RedisClient,
	responseHelper *helpers.ResponseHelper,
) *HealthHandler {
	return &HealthHandler{
		log:            log,
		dbClient:       db,
		redisClient:    redis,
		responseHelper: responseHelper,
	}
}

func (h *HealthHandler) GetHealthHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Server is running"
	h.responseHelper.SendSuccessResponse(w, msg, nil)
}

func (h *HealthHandler) GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	dbStatus := "up"
	if err := h.dbClient.HeathCheck(); err != nil {
		dbStatus = "down"
	}
	data["db"] = dbStatus

	redisStatus := "up"
	if err := h.redisClient.HealthCheck(); err != nil {
		redisStatus = "down"
	}
	data["redis"] = redisStatus

	h.responseHelper.SendSuccessResponse(w, "status update", data)
}
