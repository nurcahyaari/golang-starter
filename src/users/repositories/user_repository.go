package repositories

import (
	"golang-starter/infrastructures/db"
	"golang-starter/src/users/entities"
)

type UserRepository interface {
	FindAll() []entities.Users
	FindByID(id uint) entities.Users
	FindByEmail(email string) entities.Users
}

type userRepository struct {
	DB db.MysqlDB
}

func NewUserRepository(DB db.MysqlDB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) FindAll() []entities.Users {
	var users []entities.Users
	u.DB.Query().Find(&users)

	return users
}

func (u *userRepository) FindByID(id uint) entities.Users {
	var user entities.Users
	u.DB.Query().First(&user, id)

	return user
}

func (u *userRepository) FindByEmail(email string) entities.Users {
	var user entities.Users
	u.DB.Query().Where("email = ?", email).First(&user)

	return user
}
