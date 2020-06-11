package models

type Products struct {
	ProductID   int    `gorm:"column:product_id;primary_key" json:"product_id"`
	Name        string `gorm:"column:name" json:"name"`
	Price       string `gorm:"column:price" json:"price"`
	Description string `gorm:"column:description" json:"description"`
}

func (t *Products) TableName() string {
	return "products"
}
