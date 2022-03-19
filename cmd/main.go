package main

import (
	"context"
	"github.com/hiteshrepo/golang-jwt-auth/internal/app/di"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app, err := di.InitializeApp(context.Background())
	check(err)

	app.Start(check)

	<-interrupt()

	app.Shutdown()

	log.Default().Println("Server has stopped.")
}

func interrupt() chan os.Signal {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	return interrupt
}

func check(err error) {
	if err != nil {
		log.Default().Fatal("Could not start the server", err)
	}
}
