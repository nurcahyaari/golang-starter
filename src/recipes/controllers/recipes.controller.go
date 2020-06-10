package controllers

import (
	"golang-starter/infrastructure/db"
	"golang-starter/src/recipes/models"

	"github.com/gofiber/fiber"
)

func GetProducts(ctx *fiber.Ctx) {
	var products []models.Products
	db.Query().Find(&products)

	ctx.JSON(products)
}
