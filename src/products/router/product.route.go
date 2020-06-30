package router

import (
	"golang-starter/infrastructure/db/onlinedb"
	"golang-starter/infrastructure/middleware"
	"golang-starter/src/products/controllers"
	"golang-starter/src/products/repositories"
	"golang-starter/src/products/services"

	"github.com/gofiber/fiber"
)

func RecipesRoute(app *fiber.App) {
	db := onlinedb.Load()
	productRepository := repositories.ProvideProductRepostiory(db)
	productService := services.ProvideProductService(productRepository)
	productController := controllers.ProvideProductController(productService)

	app.Get("/products", middleware.JwtVerifyToken, productController.GetProducts)
}
