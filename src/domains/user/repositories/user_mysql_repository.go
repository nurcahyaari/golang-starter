package repositories

import (
	"golang-starter/infrastructures/db"
	"golang-starter/src/domains/user/entities"
)

type UserMysqlRepository interface {
	FindAll() []entities.Users
	FindByID(id uint) entities.Users
	FindByEmail(email string) entities.Users
}

type UserMysqlRepositoryImpl struct {
	DB *db.MysqlImpl
}

func NewUserMysqlRepository(DB *db.MysqlImpl) *UserMysqlRepositoryImpl {
	return &UserMysqlRepositoryImpl{
		DB: DB,
	}
}

func (u UserMysqlRepositoryImpl) FindAll() []entities.Users {
	var users []entities.Users
	u.DB.DB().Find(&users)

	return users
}

func (u UserMysqlRepositoryImpl) FindByID(id uint) entities.Users {
	var user entities.Users
	u.DB.DB().First(&user, id)

	return user
}

func (u UserMysqlRepositoryImpl) FindByEmail(email string) entities.Users {
	var user entities.Users
	u.DB.DB().Where("email = ?", email).First(&user)

	return user
}
