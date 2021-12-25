package router

import (
	"golang-starter/src/handlers/http"

	"github.com/gofiber/fiber/v2"
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

func (h *HttpRouterImpl) Router(r *fiber.App) {
	h.handlers.Router(r)
}
