package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hiteshrepo/golang-jwt-auth/internal/pkg/config"
	"log"
)

type App struct {
	GinEngine    *gin.Engine
	ServerConfig config.ServerConfig
}

func (a *App) Start(check func(err error)) {
	log.Default().Println("Starting server")
	go func() {
		if err := a.GinEngine.Run(fmt.Sprintf("%s:%d", a.ServerConfig.Host, a.ServerConfig.Port)); err != nil {
			check(err)
		}
	}()
}

func (a *App) Shutdown() {}