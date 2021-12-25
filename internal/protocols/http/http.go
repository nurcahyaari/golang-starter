package http

import (
	"fmt"
	"golang-starter/config"
	"golang-starter/internal/protocols/http/router"

	"github.com/gofiber/fiber/v2"
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

func (p *HttpImpl) setupRouter(app *fiber.App) {
	p.HttpRouter.Router(app)
}

func (p *HttpImpl) Listen() {

	app := fiber.New()

	p.setupRouter(app)

	app.Listen(fmt.Sprintf(":%d", config.Get().Application.Port))
}
