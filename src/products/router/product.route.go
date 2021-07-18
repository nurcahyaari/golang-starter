package router

import (
	"golang-starter/internal/db"
	"golang-starter/internal/middleware"
	"golang-starter/src/products/controllers"
	"golang-starter/src/products/repositories"
	"golang-starter/src/products/services"

	"github.com/gofiber/fiber/v2"
)

func RecipesRoute(app *fiber.App) {
	db := db.NewMysqlClient()
	productRepository := repositories.ProvideProductRepostiory(db)
	productService := services.ProvideProductService(productRepository)
	productController := controllers.ProvideProductController(productService)

	app.Get("/products", middleware.JwtVerifyToken, productController.GetProducts)
}
