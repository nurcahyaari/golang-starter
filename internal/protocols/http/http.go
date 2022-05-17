package http

import (
	"context"
	"fmt"
	"golang-starter/config"
	"golang-starter/internal/protocols/http/router"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
)

type HttpImpl struct {
	HttpRouter *router.HttpRouterImpl
}

func NewHttpProtocol(
	HttpRouter *router.HttpRouterImpl,
) *HttpImpl {
	return &HttpImpl{
		HttpRouter: HttpRouter,
	}
}

func (p *HttpImpl) setupRouter(app *chi.Mux) {
	p.HttpRouter.Router(app)
}

func (p *HttpImpl) Listen(ctx context.Context) {

	app := chi.NewRouter()

	p.setupRouter(app)

	serverPort := fmt.Sprintf(":%d", config.Get().Application.Port)
	httpserver := &http.Server{
		Addr:    serverPort,
		Handler: app,
	}

	go func() {
		log.Info().Msgf("Server started on Port %s ", serverPort)
		httpserver.ListenAndServe()
	}()

	<-ctx.Done()

	log.Info().Msg("server stopped")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer func() {
		cancel()
	}()

	if err := httpserver.Shutdown(ctxShutdown); err != nil {
		log.Err(err).Msgf("server Shutdown Failed:%+s", err)
		if err == http.ErrServerClosed {
			err = nil
		}
	}

	log.Info().Msg("server exited properly")
}
