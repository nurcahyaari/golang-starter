package repositories

import (
	localdb "golang-starter/infrastructures/local_db"
	"golang-starter/src/modules/user/entities"
)

type UserScribleRepository interface {
	FindUserRefreshToken(userID string) (entities.UserRefreshToken, error)
}

type UserScribleRepositoryImpl struct {
	scribleDB *localdb.ScribleImpl
}

func NewUserScribleRepository(scribleDB *localdb.ScribleImpl) *UserScribleRepositoryImpl {
	return &UserScribleRepositoryImpl{
		scribleDB: scribleDB,
	}
}

func (c UserScribleRepositoryImpl) FindUserRefreshToken(userID string) (entities.UserRefreshToken, error) {
	var userRefreshToken entities.UserRefreshToken
	err := c.scribleDB.DB().Read("refresh_token", userID, &userRefreshToken)
	if err != nil {
		return entities.UserRefreshToken{}, err
	}
	return userRefreshToken, nil
}
