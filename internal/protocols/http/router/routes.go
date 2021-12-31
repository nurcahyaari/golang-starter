package router

import (
	"golang-starter/src/handlers/http"

	"github.com/go-chi/chi/v5"
)

type HttpRouterImpl struct {
	handlers *http.HttpHandlerImpl
}

func NewHttpRoute(
	handlers *http.HttpHandlerImpl,
) *HttpRouterImpl {
	return &HttpRouterImpl{
		handlers: handlers,
	}
}

func (h *HttpRouterImpl) Router(r *chi.Mux) {
	h.handlers.Router(r)
}
