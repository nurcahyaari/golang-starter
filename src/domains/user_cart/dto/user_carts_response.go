package dto

type UserCartsResponseBody struct {
	CartID    int   `json:"cart_id"`
	UserID    int   `json:"user_id"`
	ProductID int   `json:"product_id"`
	Qty       int   `json:"qty"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
