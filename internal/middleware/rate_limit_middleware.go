package middleware

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/conceptcodes/eth-indexer-go/config"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type RateLimitRequestMiddleware struct {
	log *zerolog.Logger
	rdb *redis.Client
	cfg *config.Config
}

func NewRateLimitRequestMiddleware(log *zerolog.Logger, rdb *redis.Client, cfg *config.Config) *RateLimitRequestMiddleware {
	return &RateLimitRequestMiddleware{log: log, rdb: rdb, cfg: cfg}
}

func (m *RateLimitRequestMiddleware) Start(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddress, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			m.log.Error().Err(err).Msg("Error while getting IP address")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		key := fmt.Sprintf("rate_limit:%s:%d", ipAddress, time.Now().Unix())

		limiter := redis_rate.NewLimiter(m.rdb)
		res, err := limiter.Allow(r.Context(), key, redis_rate.PerSecond(m.cfg.RateLimitCapacity))

		if err != nil {
			m.log.Error().Err(err).Msg("Error while checking rate limit")
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}

		h := w.Header()
		h.Set("RateLimit-Remaining", strconv.Itoa(res.Remaining))
		h.Set("RateLimit-Limit", strconv.Itoa(m.cfg.RateLimitCapacity))

		if res.Allowed == 0 {
			seconds := int(res.RetryAfter / time.Second)
			h.Set("RateLimit-RetryAfter", strconv.Itoa(seconds))
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		m.log.Debug().Int("remaining", res.Remaining).Msg("Rate limit remaining")
		next.ServeHTTP(w, r)
	})
}
