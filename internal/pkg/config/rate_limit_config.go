package config

import "time"

type RateLimiterConfig struct {
	Enabled         bool  `mapstructure:"rateLimitEnabled"`
	CleanupInterval int   `mapstructure:"rateLimitCleanupIntervalInSeconds"`
	Limit           int64 `mapstructure:"rateLimitCount"`
	Period          int   `mapstructure:"rateLimitPeriodInMinutes"`
	WaitPeriod      int   `mapstructure:"rateLimitWaitPeriodInSeconds"`
}

func (r *RateLimiterConfig) GetLimit() int64 {
	if r.Limit == 0 {
		return defaultLimit
	}
	return r.Limit
}

func (r *RateLimiterConfig) GetPeriod() time.Duration {
	if r.Period == 0 {
		return defaultPeriod
	}
	return time.Duration(r.Period) * time.Minute
}

func (r *RateLimiterConfig) GetCleanupInterval() time.Duration {
	if r.CleanupInterval == 0 {
		return defaultCleanupInterval
	}
	return time.Duration(r.CleanupInterval) * time.Second
}

func (r *RateLimiterConfig) GetWaitPeriod() time.Duration {
	if r.WaitPeriod == 0 {
		return defaultWaitPeriod
	}
	return time.Duration(r.WaitPeriod) * time.Second
}
