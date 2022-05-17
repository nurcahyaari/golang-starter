package main

import (
	"context"
	"golang-starter/internal/graceful"
	"golang-starter/internal/logger"
	"time"
)

//go:generate go run github.com/google/wire/cmd/wire
//go:generate go run github.com/swaggo/swag/cmd/swag init

func main() {
	// you didn't define port in env file
	// the default port is random from fiber
	// init log
	logger.InitLogger()

	initProtocol := InitHttpProtocol()

	graceful.GracefulShutdown(
		context.TODO(),
		5*time.Second,
		map[string]graceful.Operation{
			"http": func(ctx context.Context) error {
				return initProtocol.Shutdown(ctx)
			},
		},
	)

	initProtocol.Listen()
}
