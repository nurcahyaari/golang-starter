package controllers

import (
	"golang-starter/src/products/models"
	"golang-starter/src/products/services"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	GetProducts(ctx *fiber.Ctx) error
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

func (services *productController) GetProducts(ctx *fiber.Ctx) error {
	var products []models.Products
	// get all products
	products = services.ProductService.GetProducts()

	return ctx.JSON(products)
}
