package repositories

import (
	"golang-starter/infrastructures/localdb"
	"golang-starter/src/modules/user/entities"
)

type UserScribleRepository interface {
	FindUserRefreshToken(userId string) (entities.UserRefreshToken, error)
}

type UserScribleRepositoryImpl struct {
	scribleDB *localdb.ScribleImpl
}

func NewUserScribleRepository(scribleDB *localdb.ScribleImpl) *UserScribleRepositoryImpl {
	return &UserScribleRepositoryImpl{
		scribleDB: scribleDB,
	}
}

func (c UserScribleRepositoryImpl) FindUserRefreshToken(userId string) (entities.UserRefreshToken, error) {
	var userRefreshToken entities.UserRefreshToken
	err := c.scribleDB.DB().Read("refresh_token", userId, &userRefreshToken)
	if err != nil {
		return entities.UserRefreshToken{}, err
	}
	return userRefreshToken, nil
}
