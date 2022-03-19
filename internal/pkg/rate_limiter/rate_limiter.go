package rate_limiter

import (
	"fmt"
	"github.com/hiteshrepo/golang-jwt-auth/internal/pkg/config"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"log"
)

func ProvideInMemoryLimiter(rateLimitCfg *config.RateLimiterConfig) (*limiter.Limiter, error) {
	if !rateLimitCfg.Enabled {
		log.Default().Println("rate limiting feature is disabled")
		return nil, nil
	}

	log.Default().Println("initializing in-memory rate limiter")

	opts := limiter.StoreOptions{
		Prefix:          limiter.DefaultPrefix,
		CleanUpInterval: rateLimitCfg.GetCleanupInterval(),
	}
	return NewLimiter(rateLimitCfg, memory.NewStoreWithOptions(opts)), nil
}

func NewLimiter(cfg *config.RateLimiterConfig, store limiter.Store) *limiter.Limiter {
	limit := cfg.GetLimit()
	period := cfg.GetPeriod()
	rate := limiter.Rate{
		Limit:  limit,
		Period: period,
	}

	log.Default().Println(fmt.Sprintf("rate limiter initialized with limit of %v request per %v minutes", limit, period.Minutes()))
	return limiter.New(store, rate)
}
