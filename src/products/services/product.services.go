package services

import (
	"golang-starter/src/products/models"
	"golang-starter/src/products/repositories"
)

type ProductService interface {
	GetProducts() []models.Products
}

type productService struct {
	ProductRepository repositories.ProductRepository
}

func ProvideProductService(
	productRepository repositories.ProductRepository,
) ProductService {
	return &productService{
		ProductRepository: productRepository,
	}
}

func (repo *productService) GetProducts() []models.Products {
	// var products []models.Products
	return repo.ProductRepository.FindAll()
}
