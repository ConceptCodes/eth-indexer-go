package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/conceptcodes/eth-indexer-go/config"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var (
	once sync.Once
)

type RedisClient struct {
	client *redis.Client
	cfg    *config.Config
	ctx    context.Context
	log    *zerolog.Logger
}

func NewRedisClient(cfg *config.Config, log *zerolog.Logger, ctx context.Context) *RedisClient {
	return &RedisClient{
		cfg: cfg,
		log: log,
		ctx: ctx,
	}
}

func (r *RedisClient) Connect() {
	once.Do(func() {
		r.log.Debug().Msg("Connecting to redis")

		r.client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", r.cfg.RedisHost, r.cfg.RedisPort),
			Password: r.cfg.RedisPassword,
			DB:       r.cfg.RedisDB,
		})

		_, err := r.client.Ping(r.ctx).Result()
		if err != nil {
			r.log.Error().Err(err).Msg("Error while connecting to redis")
		}
	})
}

func (r *RedisClient) GetClient() *redis.Client {
	return r.client
}

func (r *RedisClient) HealthCheck() error {
	_, err := r.client.Ping(r.ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
