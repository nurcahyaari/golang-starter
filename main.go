package main

import (
	"golang-starter/internal/logger"
)

//go:generate go run github.com/google/wire/cmd/wire

func main() {
	// you didn't define port in env file
	// the default port is random from fiber
	// fmt.Println(config.Get().Application.Key.Rsa.Private)
	// fmt.Println(config.Get().Application.Key.Rsa.Public)
	// init log
	logger.InitLogger()

	InitHttpProtocol().Listen()
}
