//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"github.com/hiteshrepo/golang-jwt-auth/internal/app"
	"github.com/hiteshrepo/golang-jwt-auth/internal/app/handler"
	"github.com/hiteshrepo/golang-jwt-auth/internal/app/router"
	"github.com/hiteshrepo/golang-jwt-auth/internal/pkg/config"
	"github.com/hiteshrepo/golang-jwt-auth/internal/pkg/rate_limiter"
	"github.com/google/wire"

)

var configSet = wire.NewSet(
	config.ProvideAppConfig,
	wire.FieldsOf(new(*config.AppConfig), "ServerConfig"),
	wire.FieldsOf(new(*config.AppConfig), "RateLimiterConfig"),
)

var handlerSet = wire.NewSet(
	handler.ProvideAuthHandler,
	handler.ProvideUserHandler,
)

var repoSet = wire.NewSet()

var serviceSet = wire.NewSet()

var rateLimiterSet = wire.NewSet(
	rate_limiter.ProvideInMemoryLimiter,
)

func InitializeApp(ctx context.Context) (*app.App, error) {
	wire.Build(
		configSet,
		handlerSet,
		router.ProvideRouter,

		wire.Struct(new(app.App), "*"),
	)
	return &app.App{}, nil
}
