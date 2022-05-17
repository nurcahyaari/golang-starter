package main

import (
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

	ctx := graceful.GracefulShutdown()

	InitHttpProtocol().Listen(ctx)
}
