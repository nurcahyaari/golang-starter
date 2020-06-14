package controllers

import (
	"golang-starter/src/users/dto"
	"golang-starter/src/users/services"
	"log"

	"github.com/gofiber/fiber"
)

type UserController interface {
	Login(ctx *fiber.Ctx)
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

func (service *userController) Login(ctx *fiber.Ctx) {
	userDTO := new(dto.User)

	if err := ctx.BodyParser(userDTO); err != nil {
		log.Fatal(err)
	}

	res := service.UserService.Login(userDTO)

	ctx.JSON(res)
}
