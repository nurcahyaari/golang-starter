// Code generated by "repogen"; DO NOT EDIT.
package entities

import (
	"github.com/guregu/null"
)

type Products struct {
	ProductId           int32    `db:"product_id"`
	ProductCategoryFkid null.Int `db:"product_category_fkid"`
	AdminFkid           null.Int `db:"admin_fkid"`
	Name                string   `db:"name"`
	Price               int32    `db:"price"`
	Description         string   `db:"description"`
	Qty                 int32    `db:"qty"`
	Image               string   `db:"image"`
	Label               string   `db:"label"`
}

type ProductsList []*Products