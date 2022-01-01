package dto

type ProductsResponseBody struct {
	ProductID   int    `json:"product_id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
}
