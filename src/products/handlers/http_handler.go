package handlers

import (
	"golang-starter/internal/web"
	"golang-starter/src/products/dto"
	"golang-starter/src/products/entities"
	"golang-starter/src/products/services"
	"net/http"

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
	var products []entities.Products
	// get all products
	products = services.ProductService.GetProducts()

	productsResponse := dto.ParseFromEntities(products)

	return web.JsonResponse(ctx, http.StatusOK, "", productsResponse)
}
