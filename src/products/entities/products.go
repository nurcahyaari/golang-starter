package entities

type Products struct {
	ProductID   int    `gorm:"column:product_id;primary_key"`
	Name        string `gorm:"column:name"`
	Price       string `gorm:"column:price"`
	Description string `gorm:"column:description"`
}

func (t *Products) TableName() string {
	return "products"
}
