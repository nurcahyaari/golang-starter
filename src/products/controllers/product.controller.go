package controllers

import (
	"golang-starter/src/products/models"
	"golang-starter/src/products/services"

	"github.com/gofiber/fiber"
)

type ProductController interface {
	GetProducts(ctx *fiber.Ctx)
}

type productController struct {
	ProductService services.ProductService
}

func ProvideProductController(
	productService services.ProductService,
) ProductController {
	return &productController{
		ProductService: productService,
	}
}

func (services *productController) GetProducts(ctx *fiber.Ctx) {
	var products []models.Products
	// get all products
	products = services.ProductService.GetProducts()

	ctx.JSON(products)
}
