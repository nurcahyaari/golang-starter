package services

import (
	"golang-starter/internal/utils/auth"

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
	UserRepository repositories.UserRepository
	jwtAuth        auth.JwtTokenInterface
}

func NewUserService(userRepository repositories.UserRepository, jwtAuth auth.JwtTokenInterface) UserService {
	return &userService{
		UserRepository: userRepository,
		jwtAuth:        jwtAuth,
	}
}

func (c *userService) FindByID(id uint) entities.Users {
	return c.UserRepository.FindByID(id)
}

func (c *userService) Login(data *dto.UserRequestLoginBody) (dto.UserTokenResponseBody, error) {

	var user entities.Users
	user = c.UserRepository.FindByEmail(data.Email)

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
	userToken := c.jwtAuth.Sign(jwt.MapClaims{
		"id": userID,
	})

	token := dto.UserTokenResponseBody(userToken)

	return token, nil
}
