package services

import "golang-starter/src/domains/product/repositories"

//go:generate go run github.com/sog01/repogen/cmd/repogen -module golang-starter -destination ../ -envFile .env -envPrefix DB -tables products -modelPackage entities -repositoryPackage repositories

type ProductService interface {
	// GetProducts() []entities.Products
	// GetProductByProductID(productID int) entities.Products
	// CreateNewProduct(data dto.ProductsRequestBody) (entities.Products, error)
	// DeleteProduct(productID int) error
}

type ProductServiceImpl struct {
	ProductRepository repositories.Repositories
}

func NewProductService(
	productRepository repositories.Repositories,
) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
	}
}

// func (repo ProductServiceImpl) GetProducts() []entities.Products {
// 	return repo.ProductRepository.FindAll()
// }

// func (repo ProductServiceImpl) GetProductByProductID(productID int) entities.Products {
// 	return repo.ProductRepository.FindByID(uint(productID))
// }

// func (repo ProductServiceImpl) CreateNewProduct(data dto.ProductsRequestBody) (entities.Products, error) {
// 	product := entities.Products{
// 		Name:        data.Name,
// 		Price:       data.Price,
// 		Description: data.Description,
// 		Qty:         data.Qty,
// 	}

// 	err := repo.ProductRepository.Save(&product)
// 	if err != nil {
// 		return entities.Products{}, err
// 	}

// 	return product, nil
// }

// func (repo ProductServiceImpl) DeleteProduct(productID int) error {

// 	product := repo.GetProductByProductID(productID)
// 	if product.ProductID == 0 {
// 		return sql.ErrNoRows
// 	}

// 	return repo.ProductRepository.Delete(product)
// }
