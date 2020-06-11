package repositories

import (
	"golang-starter/infrastructure/db"
	"golang-starter/src/products/models"
)

type ProductRepository interface {
	FindAll() []models.Products
	FindByID(id uint) models.Products
	Save(product models.Products) models.Products
	Delete(product models.Products)
}

type productRepository struct {
	DB db.Database
}

func ProvideProductRepostiory(DB db.Database) ProductRepository {
	return &productRepository{
		DB: DB,
	}
}

func (p *productRepository) FindAll() []models.Products {
	var products []models.Products
	p.DB.Query().Find(&products)

	return products
}

func (p *productRepository) FindByID(id uint) models.Products {
	var product models.Products
	p.DB.Query().First(&product, id)

	return product
}

func (p *productRepository) Save(product models.Products) models.Products {
	p.DB.Query().Save(&product)

	return product
}

func (p *productRepository) Delete(product models.Products) {
	p.DB.Query().Delete(&product)
}
