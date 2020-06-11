package repositories

import (
	"golang-starter/infrastructure/db"
	"golang-starter/src/products/models"
)

func FindAll() []models.Products {
	var products []models.Products
	db.Query().Find(&products)

	return products
}

func FindByID(id uint) models.Products {
	var product models.Products
	db.Query().First(&product, id)

	return product
}

func Save(product models.Products) models.Products {
	db.Query().Save(&product)

	return product
}

func Delete(product models.Products) {
	db.Query().Delete(&product)
}
