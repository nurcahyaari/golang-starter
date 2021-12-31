package http

import (
	"fmt"
	"golang-starter/config"
	"golang-starter/internal/protocols/http/router"
	"net/http"

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

func (p *HttpImpl) Listen() {

	app := chi.NewRouter()

	p.setupRouter(app)

	http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Application.Port), app)
}
