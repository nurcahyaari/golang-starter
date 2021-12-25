package dto

type UserCartsRequestBody struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Qty       int `json:"qty"`
}
