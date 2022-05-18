package graceful

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type Operation func(ctx context.Context) error

func GracefulShutdown(ctx context.Context, timeout time.Duration, operations map[string]Operation) {
	if len(operations) == 0 {
		return
	}

	wait := make(chan struct{})
	go func() {
		signalchan := make(chan os.Signal, 1)
		signal.Notify(signalchan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		oscall := <-signalchan

		timeAfterExecuted := time.AfterFunc(timeout, func() {
			log.Warn().Msg("Force shutdown")
			os.Exit(0)
		})
		defer timeAfterExecuted.Stop()

		wg := sync.WaitGroup{}
		wg.Add(len(operations))
		for k, op := range operations {
			go func(k string, op Operation) {
				defer wg.Done()
				log.Warn().Msgf("Shutdown %s", k)
				err := op(ctx)
				if err != nil {
					log.Err(err).Msg("Error when stop server")
				}
			}(k, op)
		}
		wg.Wait()

		log.Warn().Msgf("system call:%+v", oscall)
		close(wait)
	}()
}
