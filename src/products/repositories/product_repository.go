package repositories

import (
	"golang-starter/infrastructures/db"
	"golang-starter/src/products/entities"
)

type ProductRepository interface {
	FindAll() []entities.Products
	FindByID(id uint) entities.Products
	Save(product entities.Products) entities.Products
	Delete(product entities.Products)
}

type productRepository struct {
	DB db.MysqlDB
}

func NewProductRepostiory(DB db.MysqlDB) ProductRepository {
	return &productRepository{
		DB: DB,
	}
}

func (p *productRepository) FindAll() []entities.Products {
	var products []entities.Products
	p.DB.Query().Find(&products)

	return products
}

func (p *productRepository) FindByID(id uint) entities.Products {
	var product entities.Products
	p.DB.Query().First(&product, id)

	return product
}

func (p *productRepository) Save(product entities.Products) entities.Products {
	p.DB.Query().Save(&product)

	return product
}

func (p *productRepository) Delete(product entities.Products) {
	p.DB.Query().Delete(&product)
}
