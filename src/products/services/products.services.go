package services

import (
	"golang-starter/src/products/models"
	"golang-starter/src/products/repositories"
)

func GetProducts() []models.Products {
	// var products []models.Products
	return repositories.FindAll()
}
