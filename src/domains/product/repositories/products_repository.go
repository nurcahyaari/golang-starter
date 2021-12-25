package repositories

import (
	"golang-starter/infrastructures/db"
	"golang-starter/src/domains/product/entities"
)

type ProductRepository interface {
	FindAll() []entities.Products
	FindByID(id uint) entities.Products
	Save(product *entities.Products) error
	Delete(product entities.Products) error
}

type ProductRepositoryImpl struct {
	DB *db.MysqlImpl
}

func NewProductRepostiory(DB *db.MysqlImpl) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		DB: DB,
	}
}

func (p ProductRepositoryImpl) FindAll() []entities.Products {
	var products []entities.Products
	p.DB.DB().Find(&products)

	return products
}

func (p ProductRepositoryImpl) FindByID(id uint) entities.Products {
	var product entities.Products
	p.DB.DB().First(&product, id)

	return product
}

func (p ProductRepositoryImpl) Save(product *entities.Products) error {
	err := p.DB.DB().Save(&product).Error

	if err != nil {
		return err
	}

	return nil
}

func (p ProductRepositoryImpl) Delete(product entities.Products) error {
	err := p.DB.DB().Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
