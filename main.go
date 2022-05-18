package main

import (
	"context"
	"golang-starter/config"
	"golang-starter/internal/graceful"
	"golang-starter/internal/logger"
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
		config.Get().Application.Graceful.MaxSecond,
		map[string]graceful.Operation{
			"http": func(ctx context.Context) error {
				return initProtocol.Shutdown(ctx)
			},
		},
	)

	initProtocol.Listen()
}
