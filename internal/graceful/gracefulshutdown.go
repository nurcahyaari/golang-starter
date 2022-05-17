package graceful

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type Operation func(ctx context.Context) error

func GracefulShutdown(ctx context.Context, timeout time.Duration, operations map[string]Operation) {
	if len(operations) == 0 {
		return
	}

	go func() {
		signalchan := make(chan os.Signal, 1)
		signal.Notify(signalchan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		oscall := <-signalchan

		for k, op := range operations {
			go func(k string, op Operation) {
				log.Warn().Msgf("Shutdown %s", k)
				err := op(ctx)
				if err != nil {
					log.Err(err).Msg("Error when stop server")
					// return
				}
				log.Warn().Msgf("%s server stopped", k)
			}(k, op)
		}

		log.Warn().Msgf("system call:%+v", oscall)

		time.Sleep(timeout)
		os.Exit(0)
	}()
}
