package dto

import "golang-starter/src/domains/product/entities"

type ProductsResponseBody struct {
	ProductID   int    `json:"product_id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
}

func ParseFromEntity(data entities.Products) ProductsResponseBody {
	return ProductsResponseBody{
		ProductID:   data.ProductID,
		Name:        data.Name,
		Price:       data.Price,
		Description: data.Description,
		Qty:         data.Qty,
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
			Qty:         d.Qty,
		}
		productsResponseBodies = append(productsResponseBodies, productsResponseBody)
	}

	return productsResponseBodies
}
