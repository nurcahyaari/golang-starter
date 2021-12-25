package http

import (
	httpresponse "golang-starter/internal/protocols/http/response"
	"golang-starter/src/domains/user/dto"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h HttpHandlerImpl) Login(ctx *fiber.Ctx) error {
	userData := new(dto.UserRequestLoginBody)

	if err := ctx.BodyParser(userData); err != nil {
		log.Fatal(err)
	}

	res, err := h.UserService.Login(userData)
	if err != nil {
		return httpresponse.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return httpresponse.JsonResponse(ctx, http.StatusOK, "", res)
}

func (h HttpHandlerImpl) Refresh(ctx *fiber.Ctx) error {
	userID := ctx.Get("userID")

	res, err := h.UserService.RefreshToken(userID)

	if err != nil {
		return httpresponse.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return httpresponse.JsonResponse(ctx, http.StatusOK, "", res)
}
