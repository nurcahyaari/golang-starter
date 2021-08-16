package services

import (
	"database/sql"
	"golang-starter/internal/utils/auth"
	"time"

	"golang-starter/src/users/dto"
	"golang-starter/src/users/entities"
	"golang-starter/src/users/repositories"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindByID(id uint) entities.Users
	Login(data *dto.UserRequestLoginBody) (dto.UserTokenResponseBody, error)
	RefreshToken(userID string) (dto.UserTokenResponseBody, error)
}

type userService struct {
	userMysqlRepository   repositories.UserMysqlRepositoryInterface
	userScribleRepository repositories.UserScribleRepositoryInterface
	jwtAuth               auth.JwtTokenInterface
}

func NewUserService(
	userMysqlRepository repositories.UserMysqlRepositoryInterface,
	jwtAuth auth.JwtTokenInterface,
	userScribleRepository repositories.UserScribleRepositoryInterface,
) UserService {
	return &userService{
		userScribleRepository: userScribleRepository,
		userMysqlRepository:   userMysqlRepository,
		jwtAuth:               jwtAuth,
	}
}

func (c *userService) FindByID(id uint) entities.Users {
	return c.userMysqlRepository.FindByID(id)
}

func (c *userService) Login(data *dto.UserRequestLoginBody) (dto.UserTokenResponseBody, error) {

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

func (c *userService) RefreshToken(userID string) (dto.UserTokenResponseBody, error) {
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
