package controllers

import (
	"golang-starter/src/users/dto"
	"golang-starter/src/users/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Login(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
}

type userController struct {
	UserService services.UserService
}

func ProvideUserController(
	userService services.UserService,
) UserController {
	return &userController{
		UserService: userService,
	}
}

func (service *userController) Login(ctx *fiber.Ctx) error {
	userDTO := new(dto.User)

	if err := ctx.BodyParser(userDTO); err != nil {
		log.Fatal(err)
	}

	res := service.UserService.Login(userDTO)

	return ctx.JSON(res)
}

func (service *userController) Refresh(ctx *fiber.Ctx) error {
	userID := ctx.Get("userID")

	res := service.UserService.RefreshToken(userID)

	return ctx.JSON(res)
}
