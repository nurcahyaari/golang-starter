package services

import (
	"database/sql"
	"golang-starter/internal/utils/auth"
	"golang-starter/src/domains/user/dto"
	"golang-starter/src/domains/user/entities"
	"golang-starter/src/domains/user/repositories"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindByID(id uint) entities.Users
	Login(data *dto.UserRequestLoginBody) (dto.UserTokenResponseBody, error)
	RefreshToken(userID string) (dto.UserTokenResponseBody, error)
}

type UserServiceImpl struct {
	userMysqlRepository   repositories.UserMysqlRepository
	userScribleRepository repositories.UserScribleRepository
	jwtAuth               auth.JwtToken
}

func NewUserService(
	jwtAuth auth.JwtToken,
	userMysqlRepository repositories.UserMysqlRepository,
	userScribleRepository repositories.UserScribleRepository,
) *UserServiceImpl {
	return &UserServiceImpl{
		userScribleRepository: userScribleRepository,
		userMysqlRepository:   userMysqlRepository,
		jwtAuth:               jwtAuth,
	}
}

func (c UserServiceImpl) FindByID(id uint) entities.Users {
	return c.userMysqlRepository.FindByID(id)
}

func (c UserServiceImpl) Login(data *dto.UserRequestLoginBody) (dto.UserTokenResponseBody, error) {

	user := c.userMysqlRepository.FindByEmail(data.Email)
	if user.UserID == 0 {
		return dto.UserTokenResponseBody{}, sql.ErrNoRows
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return dto.UserTokenResponseBody{}, err
	}

	userToken := c.jwtAuth.Sign(jwt.MapClaims{
		"id": user.UserID,
	})

	token := dto.UserTokenResponseBody(userToken)

	return token, nil
}

func (c UserServiceImpl) RefreshToken(userID string) (dto.UserTokenResponseBody, error) {
	refreshToken, err := c.userScribleRepository.FindUserRefreshToken(userID)
	if err != nil {
		return dto.UserTokenResponseBody{}, err
	}

	if refreshToken.Expired < time.Now().Unix() {
		return dto.UserTokenResponseBody{}, err
	}

	userToken := c.jwtAuth.Sign(jwt.MapClaims{
		"id": userID,
	})

	token := dto.UserTokenResponseBody(userToken)

	return token, nil
}
