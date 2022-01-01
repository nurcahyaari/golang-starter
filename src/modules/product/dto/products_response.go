package dto

import (
	"golang-starter/src/modules/product/entities"
	"strconv"
)

type ProductsResponse struct {
	ProductID   int    `json:"product_id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
}

func CreateProductsResponse(product entities.Products) ProductsResponse {
	return ProductsResponse{
		ProductID:   int(product.ProductId),
		Name:        product.Name,
		Price:       strconv.Itoa(int(product.Price)),
		Description: product.Description,
		Qty:         int(product.Qty),
	}
}

type ProductsListResponse []*ProductsResponse

func CreateProductsListResponse(products entities.ProductsList) ProductsListResponse {
	productsResp := ProductsListResponse{}
	for _, p := range products {
		product := CreateProductsResponse(*p)
		productsResp = append(productsResp, &product)
	}
	return productsResp
}
