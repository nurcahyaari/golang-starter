package services

import (
	"golang-starter/infrastructures/db"
	"golang-starter/infrastructures/logger"
	productsRepositories "golang-starter/src/domains/products/repositories"
	"golang-starter/src/domains/user_cart/dto"
	"golang-starter/src/domains/user_cart/entities"
	"golang-starter/src/domains/user_cart/repositories"
	usersRepositories "golang-starter/src/domains/users/repositories"
)

type UserCartsServiceInterface interface {
	FindUserCart(userID int) ([]entities.UserCarts, error)
	AddToUserCart(data dto.UserCartsRequestBody) (entities.UserCarts, error)
}

type userCartsService struct {
	userCartsRepo repositories.UserCartRepositoryInterface
	userRepo      usersRepositories.UserMysqlRepositoryInterface
	productRepo   productsRepositories.ProductRepository
	mysqlDB       db.MysqlDB
}

func NewUserCartsService(
	userCartsRepo repositories.UserCartRepositoryInterface,
	userRepo usersRepositories.UserMysqlRepositoryInterface,
	productRepo productsRepositories.ProductRepository,
	mysqlDB db.MysqlDB,
) UserCartsServiceInterface {
	return &userCartsService{
		userCartsRepo: userCartsRepo,
		userRepo:      userRepo,
		productRepo:   productRepo,
		mysqlDB:       mysqlDB,
	}
}

func (c userCartsService) FindUserCart(userID int) ([]entities.UserCarts, error) {
	userCart, err := c.userCartsRepo.FindUserCart(userID)
	if err != nil {
		logger.Log.Errorln(err)
		return nil, err
	}
	return userCart, nil
}

func (c userCartsService) AddToUserCart(data dto.UserCartsRequestBody) (entities.UserCarts, error) {
	// tx := c.mysqlDB.DB().Begin()

	return entities.UserCarts{}, nil
}
