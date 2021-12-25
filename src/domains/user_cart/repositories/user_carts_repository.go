package repositories

import (
	"golang-starter/infrastructures/db"
	"golang-starter/src/domains/user_cart/entities"
)

type UserCartRepositoryInterface interface {
	FindUserCart(userID int) ([]entities.UserCarts, error)
	// SaveTx(tx *sql.Tx, data entities.UserCarts) error
}

type userCartRepository struct {
	mysqlDB db.MysqlDB
}

func NewUserCartRepository(mysqlDB db.MysqlDB) UserCartRepositoryInterface {
	return &userCartRepository{
		mysqlDB: mysqlDB,
	}
}

func (c userCartRepository) FindUserCart(userID int) ([]entities.UserCarts, error) {
	var userCarts []entities.UserCarts
	err := c.mysqlDB.DB().Find(&userCarts).Error
	if err != nil {
		return nil, err
	}

	return userCarts, nil

}

// func (c userCartRepository) SaveTx(tx *sql.Tx, data entities.UserCarts) error {

// }
