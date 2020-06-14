package services

import (
	"golang-starter/infrastructure/utils/auth"
	"golang-starter/infrastructure/utils/response"
	"golang-starter/src/users/dto"
	"golang-starter/src/users/models"
	"golang-starter/src/users/repositories"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindByID(id uint) models.Users
	Login(userDTO *dto.User) response.ResponseDTO
}

type userService struct {
	UserRepository repositories.UserRepository
}

func ProvideUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		UserRepository: userRepository,
	}
}

func (repo *userService) FindByID(id uint) models.Users {
	return repo.UserRepository.FindByID(id)
}

func (repo *userService) Login(userDTO *dto.User) response.ResponseDTO {
	var user models.Users
	user = repo.UserRepository.FindByEmail(userDTO.Email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.Password))

	if err != nil {
		return response.ResponseDTO{
			Message: "Username or Password incorrect",
		}
	}

	userToken := auth.Sign(jwt.MapClaims{
		"id": user.UserID,
	})

	return response.ResponseDTO{
		Data: userToken,
	}
}
