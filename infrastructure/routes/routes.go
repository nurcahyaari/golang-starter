package routes

import (
	recipeRouter "golang-starter/src/recipes/router"

	"github.com/gofiber/fiber"
)

func RegisterRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello this is my first route in go fiber")
	})

	// app.Get("/:name", func(c *fiber.Ctx) {
	// 	c.Send("Hello", c.Params("name"))
	// })

	// registering route from another modules
	recipeRouter.RecipesRoute(app)

	// handling 404 error
	app.Use(func(c *fiber.Ctx) {
		c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
}
