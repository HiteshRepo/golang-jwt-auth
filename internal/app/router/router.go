package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/hiteshrepo/golang-jwt-auth/internal/app/handler"
	"github.com/hiteshrepo/golang-jwt-auth/internal/pkg/config"
	ginTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

func ProvideRouter(appConfig *config.AppConfig, authHandler handler.AuthHandler, userHandler handler.UserHandler) (*gin.Engine, error) {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(ginTrace.Middleware(appConfig.GetServerConfig().ServiceName))

	g := r.Group("api/v1")
	{
		g.GET("/token", authHandler.FetchToken)
		g.GET("/users", userHandler.ListUsers)
		g.POST("/users/signup", authHandler.Signup)
		g.POST("/users/login", authHandler.Login)
	}

	return r, nil
}
