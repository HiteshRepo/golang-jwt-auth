package rate_limiter

import (
	"github.com/gin-gonic/gin"
	rateLimiterGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

const (
	RateLimitPrefix = "limiter"
)

func RateLimitKeyGetter() rateLimiterGin.KeyGetter {
	return mkKeyUsingClientIP()
}

func mkKeyUsingClientIP() func(c *gin.Context) string {
	return func(c *gin.Context) string {
		return prependPrefix(c.ClientIP(), RateLimitPrefix)
	}
}

func prependPrefix(value, prefix string) string {
	if prefix == "" {
		return value
	}
	return prefix + "-" + value
}
