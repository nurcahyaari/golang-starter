package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonResponse(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	resp := Response{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}

	return ctx.Status(statusCode).JSON(resp)
}
