package dto

import "golang-starter/src/products/entities"

type ProductsResponseBody struct {
	ProductID   int    `json:"product_id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

func ParseFromEntity(data entities.Products) ProductsResponseBody {
	return ProductsResponseBody{
		ProductID:   data.ProductID,
		Name:        data.Name,
		Price:       data.Price,
		Description: data.Description,
	}
}

func ParseFromEntities(data []entities.Products) []ProductsResponseBody {
	var productsResponseBodies []ProductsResponseBody

	for _, d := range data {
		productsResponseBody := ProductsResponseBody{
			ProductID:   d.ProductID,
			Name:        d.Name,
			Price:       d.Price,
			Description: d.Description,
		}
		productsResponseBodies = append(productsResponseBodies, productsResponseBody)
	}

	return productsResponseBodies
}
