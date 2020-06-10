package router

import (
	"golang-starter/src/recipes/controllers"

	"github.com/gofiber/fiber"
)

func RecipesRoute(app *fiber.App) {
	app.Get("/products", controllers.GetProducts)
}
