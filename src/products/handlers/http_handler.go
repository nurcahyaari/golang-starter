package handlers

import (
	"golang-starter/src/products/models"
	"golang-starter/src/products/services"

	"github.com/gofiber/fiber/v2"
)

type ProductHandlers interface {
	GetProducts(ctx *fiber.Ctx) error
}

type productHandlers struct {
	ProductService services.ProductService
}

func NewHttpHandler(
	productService services.ProductService,
) ProductHandlers {
	return &productHandlers{
		ProductService: productService,
	}
}

func (services *productHandlers) GetProducts(ctx *fiber.Ctx) error {
	var products []models.Products
	// get all products
	products = services.ProductService.GetProducts()

	return ctx.JSON(products)
}
