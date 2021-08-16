package repositories

import (
	"golang-starter/infrastructures/local_db"
	"golang-starter/src/users/entities"
)

type UserScribleRepositoryInterface interface {
	FindUserRefreshToken(userID string) (entities.UserRefreshToken, error)
}

type userScribleRepository struct {
	scribleDB local_db.ScribleDB
}

func NewUserScribleRepositoryInterface(scribleDB local_db.ScribleDB) UserScribleRepositoryInterface {
	return &userScribleRepository{
		scribleDB: scribleDB,
	}
}

func (c userScribleRepository) FindUserRefreshToken(userID string) (entities.UserRefreshToken, error) {
	var userRefreshToken entities.UserRefreshToken
	err := c.scribleDB.DB().Read("refresh_token", userID, &userRefreshToken)
	if err != nil {
		return entities.UserRefreshToken{}, err
	}
	return userRefreshToken, nil
}
