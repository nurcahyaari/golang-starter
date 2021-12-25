package entities

type UserCarts struct {
	CartID    int   `gorm:"column:cart_id"`
	UserID    int   `gorm:"column:user_id"`
	ProductID int   `gorm:"column:product_id"`
	Qty       int   `gorm:"column:qty"`
	CreatedAt int64 `gorm:"column:created_at"`
	UpdatedAt int64 `gorm:"column:updated_at"`
}

func (c *UserCarts) TableName() string {
	return "users_cart"
}
