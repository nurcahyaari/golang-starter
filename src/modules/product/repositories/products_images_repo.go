package repositories

import (
	"context"
	productsimagesmodel "golang-starter/src/modules/product/entities"

	"github.com/nurcahyaari/sqlabst"
)

type ProductsImagesRepo interface {
	GetProductsImagesCount(ctx context.Context) (int, error)
	GetProductsImagesById(ctx context.Context, id int64) (*productsimagesmodel.ProductsImages, error)
	GetProductsImagesList(ctx context.Context) (productsimagesmodel.ProductsImagesList, error)
	InsertProductsImagesList(ctx context.Context, productsImagesList productsimagesmodel.ProductsImagesList) (*InsertResult, error)
	InsertProductsImages(ctx context.Context, productsImages *productsimagesmodel.ProductsImages) (*InsertResult, error)
	// UpdateProductsImagesByFilter(ctx context.Context, productsImages *productsimagesmodel.ProductsImages, filter Filter, updatedFields ...ProductsImagesField) error
	// UpdateProductsImages(ctx context.Context, productsImages *productsimagesmodel.ProductsImages, productimagesid int32, updatedFields ...ProductsImagesField) error
	DeleteProductsImagesList(ctx context.Context, filter Filter) error
	DeleteProductsImages(ctx context.Context, productimagesid int32) error
}

type ProductsImagesRepoImpl struct {
	db *sqlabst.SqlAbst
}

func NewProductsImagesRepo(db *sqlabst.SqlAbst) *ProductsImagesRepoImpl {
	return &ProductsImagesRepoImpl{
		db: db,
	}
}

func (repo ProductsImagesRepoImpl) GetProductsImagesById(ctx context.Context, id int64) (*productsimagesmodel.ProductsImages, error)
