package controllers

import (
	"golang-starter/src/products/models"
	"golang-starter/src/products/services"

	"github.com/gofiber/fiber"
)

func GetProducts(ctx *fiber.Ctx) {
	var products []models.Products

	// get all products
	products = services.GetProducts()

	ctx.JSON(products)
}
