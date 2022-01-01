package dto

type ProductsRequestBody struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
}

type ProductsRequestParams struct {
	ProductID int
}
