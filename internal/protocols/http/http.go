package http

import (
	"context"
	"fmt"
	"golang-starter/config"
	"golang-starter/internal/protocols/http/router"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
)

type HttpImpl struct {
	HttpRouter *router.HttpRouterImpl
	httpServer *http.Server
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

func (p *HttpImpl) Listen() {
	app := chi.NewRouter()

	p.setupRouter(app)

	serverPort := fmt.Sprintf(":%d", config.Get().Application.Port)
	p.httpServer = &http.Server{
		Addr:    serverPort,
		Handler: app,
	}

	log.Info().Msgf("Server started on Port %s ", serverPort)
	p.httpServer.ListenAndServe()
}

func (p *HttpImpl) Shutdown(ctx context.Context) error {
	if err := p.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
