package entities

type Products struct {
	ProductID   int    `gorm:"column:product_id;primary_key"`
	Name        string `gorm:"column:name"`
	Price       string `gorm:"column:price"`
	Description string `gorm:"column:description"`
	Qty         int    `gorm:"column:qty"`
}

func (t *Products) TableName() string {
	return "products"
}
