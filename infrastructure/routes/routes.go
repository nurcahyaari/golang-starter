package routes

import (
	productRouter "golang-starter/src/products/router"
	userRouter "golang-starter/src/users/router"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello this is my first route in go fiber"))
	})

	// registering route from another modules
	productRouter.RecipesRoute(app)
	userRouter.UserRoute(app)

	// handling 404 error
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
}
