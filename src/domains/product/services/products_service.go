package services

import (
	"database/sql"
	"golang-starter/src/domains/product/dto"
	"golang-starter/src/domains/product/entities"
	"golang-starter/src/domains/product/repositories"
)

type ProductService interface {
	GetProducts() []entities.Products
	GetProductByProductID(productID int) entities.Products
	CreateNewProduct(data dto.ProductsRequestBody) (entities.Products, error)
	DeleteProduct(productID int) error
}

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(
	productRepository repositories.ProductRepository,
) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
	}
}

func (repo ProductServiceImpl) GetProducts() []entities.Products {
	return repo.ProductRepository.FindAll()
}

func (repo ProductServiceImpl) GetProductByProductID(productID int) entities.Products {
	return repo.ProductRepository.FindByID(uint(productID))
}

func (repo ProductServiceImpl) CreateNewProduct(data dto.ProductsRequestBody) (entities.Products, error) {
	product := entities.Products{
		Name:        data.Name,
		Price:       data.Price,
		Description: data.Description,
		Qty:         data.Qty,
	}

	err := repo.ProductRepository.Save(&product)
	if err != nil {
		return entities.Products{}, err
	}

	return product, nil
}

func (repo ProductServiceImpl) DeleteProduct(productID int) error {

	product := repo.GetProductByProductID(productID)
	if product.ProductID == 0 {
		return sql.ErrNoRows
	}

	return repo.ProductRepository.Delete(product)
}
