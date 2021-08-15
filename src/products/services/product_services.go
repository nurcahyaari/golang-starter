package services

import (
	"golang-starter/src/products/entities"
	"golang-starter/src/products/repositories"
)

type ProductService interface {
	GetProducts() []entities.Products
}

type productService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(
	productRepository repositories.ProductRepository,
) ProductService {
	return &productService{
		ProductRepository: productRepository,
	}
}

func (repo *productService) GetProducts() []entities.Products {
	// var products []entities.Products
	return repo.ProductRepository.FindAll()
}
