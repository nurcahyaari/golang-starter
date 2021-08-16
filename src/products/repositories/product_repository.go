package repositories

import (
	"golang-starter/infrastructures/db"
	"golang-starter/src/products/entities"
)

type ProductRepositoryInterface interface {
	FindAll() []entities.Products
	FindByID(id uint) entities.Products
	Save(product entities.Products) entities.Products
	Delete(product entities.Products)
}

type productRepository struct {
	DB db.MysqlDB
}

func NewProductRepostiory(DB db.MysqlDB) ProductRepositoryInterface {
	return &productRepository{
		DB: DB,
	}
}

func (p *productRepository) FindAll() []entities.Products {
	var products []entities.Products
	p.DB.DB().Find(&products)

	return products
}

func (p *productRepository) FindByID(id uint) entities.Products {
	var product entities.Products
	p.DB.DB().First(&product, id)

	return product
}

func (p *productRepository) Save(product entities.Products) entities.Products {
	p.DB.DB().Save(&product)

	return product
}

func (p *productRepository) Delete(product entities.Products) {
	p.DB.DB().Delete(&product)
}
