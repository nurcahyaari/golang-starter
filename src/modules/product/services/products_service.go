package services

import (
	"context"
	"golang-starter/infrastructures/db/transaction"
	"golang-starter/src/modules/product/dto"
	"golang-starter/src/modules/product/repositories"

	"github.com/rs/zerolog/log"
)

type ProductService interface {
	GetProducts(ctx context.Context) (dto.ProductsListResponse, error)
	GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error)
	CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error)
	DeleteProduct(ctx context.Context, productID int) error
}

type ProductServiceImpl struct {
	ProductRepository repositories.Repositories
	transaction       *transaction.TransactionImpl
}

func NewProductService(
	productRepository repositories.Repositories,
	transaction *transaction.TransactionImpl,
) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		transaction:       transaction,
	}
}

func (s ProductServiceImpl) GetProducts(ctx context.Context) (dto.ProductsListResponse, error) {
	productList, err := s.ProductRepository.GetProductsList(ctx)
	if err != nil {
		log.Err(err).Msg("Error fetch productList from DB")
	}
	productsResp := dto.CreateProductsListResponse(productList)
	return productsResp, nil
}

func (s ProductServiceImpl) GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error) {
	product, err := s.ProductRepository.
		FilterProducts(
			repositories.
				NewProductsFilter("AND").
				SetFilterByProductId(productID, "="),
		).
		GetProducts(ctx)
	if err != nil {
		log.Err(err).Msg("Error fetch productList from DB")
	}
	productResp := dto.CreateProductsResponse(*product)
	return productResp, nil
}

func (s ProductServiceImpl) CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error) {
	product := data.ToProductEntities()

	// start transaction
	err := s.transaction.RunWithTransaction(ctx, func() error {
		res, err := s.ProductRepository.InsertProducts(ctx, product)
		if err != nil {
			return err
		}

		lastInsertedId, err := res.LastInsertId()
		if err != nil {
			return err
		}

		productImages := data.ToProductImagesEntities(lastInsertedId)

		_, err = s.ProductRepository.InsertProductsImagesList(ctx, productImages)
		if err != nil {
			return err
		}
		return nil
	})

	return nil, err
}

func (s ProductServiceImpl) DeleteProduct(ctx context.Context, productID int) error {
	err := s.ProductRepository.DeleteProducts(ctx, int32(productID))

	if err != nil {
		log.Err(err).Msg("Error deleting product")
		return err
	}

	return nil
}
