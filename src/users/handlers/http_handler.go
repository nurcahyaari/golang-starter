package handlers

import (
	"golang-starter/internal/web"
	"golang-starter/src/users/dto"
	"golang-starter/src/users/services"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers interface {
	Login(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
}

type userHandlers struct {
	UserService services.UserService
}

func NewHttpHandler(
	userService services.UserService,
) UserHandlers {
	return &userHandlers{
		UserService: userService,
	}
}

func (service *userHandlers) Login(ctx *fiber.Ctx) error {
	userData := new(dto.UserRequestLoginBody)

	if err := ctx.BodyParser(userData); err != nil {
		log.Fatal(err)
	}

	res, err := service.UserService.Login(userData)
	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "", res)
}

func (service *userHandlers) Refresh(ctx *fiber.Ctx) error {
	userID := ctx.Get("userID")

	res, err := service.UserService.RefreshToken(userID)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "", res)
}
