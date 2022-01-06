package dto

import (
	"golang-starter/src/modules/product/entities"

	"github.com/guregu/null"
)

type ProductRequestBody struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Qty         int      `json:"qty"`
	Images      []string `json:"images"`
}

func (product ProductRequestBody) ToProductEntities() *entities.Products {
	return &entities.Products{
		Name:        product.Name,
		Description: product.Name,
		Price:       int32(product.Price),
		Qty:         int32(product.Qty),
	}
}

func (product ProductRequestBody) ToProductImagesEntities(productId int64) entities.ProductsImagesList {
	productImages := entities.ProductsImagesList{}

	for _, image := range product.Images {
		productImages = append(productImages, &entities.ProductsImages{
			ProductFkid: null.IntFrom(productId),
			Images:      image,
		})
	}

	return productImages
}

type ProductsRequestParams struct {
	ProductID int
}
