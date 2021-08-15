package dto

type ProductsRequestBody struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
