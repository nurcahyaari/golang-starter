package repositories

import (
	"golang-starter/infrastructures/db"
	"golang-starter/src/users/entities"
)

type UserMysqlRepositoryInterface interface {
	FindAll() []entities.Users
	FindByID(id uint) entities.Users
	FindByEmail(email string) entities.Users
}

type userMysqlRepository struct {
	DB db.MysqlDB
}

func NewUserMysqlRepository(DB db.MysqlDB) UserMysqlRepositoryInterface {
	return &userMysqlRepository{
		DB: DB,
	}
}

func (u *userMysqlRepository) FindAll() []entities.Users {
	var users []entities.Users
	u.DB.DB().Find(&users)

	return users
}

func (u *userMysqlRepository) FindByID(id uint) entities.Users {
	var user entities.Users
	u.DB.DB().First(&user, id)

	return user
}

func (u *userMysqlRepository) FindByEmail(email string) entities.Users {
	var user entities.Users
	u.DB.DB().Where("email = ?", email).First(&user)

	return user
}
