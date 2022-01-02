package dto

import "golang-starter/src/modules/product/entities"

type ProductRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
}

func (product ProductRequestBody) ToProductEntities() *entities.Products {
	return &entities.Products{
		Name:        product.Name,
		Description: product.Name,
		Price:       int32(product.Price),
		Qty:         int32(product.Qty),
	}
}

type ProductsRequestParams struct {
	ProductID int
}
