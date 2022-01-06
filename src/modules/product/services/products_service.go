package services

import (
	"context"
	"golang-starter/src/modules/product/dto"
	"golang-starter/src/modules/product/repositories"

	"github.com/rs/zerolog/log"
)

//go:generate go run github.com/sog01/repogen/cmd/repogen -module golang-starter -destination ../ -envFile .env -envPrefix DB -tables products -modelPackage entities -repositoryPackage repositories
//go:generate go run github.com/sog01/repogen/cmd/repogen -module golang-starter -destination ../ -envFile .env -envPrefix DB -tables products_images -modelPackage entities -repositoryPackage repositories

type ProductService interface {
	GetProducts(ctx context.Context) (dto.ProductsListResponse, error)
	GetProductByProductID(ctx context.Context, productID int) (dto.ProductsResponse, error)
	CreateNewProduct(ctx context.Context, data dto.ProductRequestBody) (*dto.ProductsResponse, error)
	DeleteProduct(ctx context.Context, productID int) error
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

	// start repository
	var err error
	tx := s.ProductRepository.StartTx()

	defer func() {
		if err != nil {
			log.Err(err).Msg("an error occured")
			tx.Rollback()
		}
	}()

	res, err := s.ProductRepository.InsertProducts(ctx, product)
	if err != nil {
		return nil, err
	}

	lastInsertedId, err := res.LastInsertId()

	productImages := data.ToProductImagesEntities(lastInsertedId)

	_, err = s.ProductRepository.InsertProductsImagesList(ctx, productImages)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s ProductServiceImpl) DeleteProduct(ctx context.Context, productID int) error {
	err := s.ProductRepository.DeleteProducts(ctx, int32(productID))

	if err != nil {
		log.Err(err).Msg("Error deleting product")
		return err
	}

	return nil
}
