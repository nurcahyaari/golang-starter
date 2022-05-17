package graceful

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
)

func GracefulShutdown() context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Info().Msgf("system call:%+v", oscall)
		cancel()
	}()

	return ctx
}
