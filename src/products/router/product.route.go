package router

import (
	"golang-starter/src/products/controllers"

	"github.com/gofiber/fiber"
)

func RecipesRoute(app *fiber.App) {
	app.Get("/products", controllers.GetProducts)
}
