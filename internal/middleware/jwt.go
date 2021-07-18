package middleware

import (
	"fmt"
	"golang-starter/internal/config"
	"golang-starter/internal/db"
	"golang-starter/internal/utils/auth"
	"golang-starter/internal/utils/response"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/gofiber/fiber/v2"
)

/*
** Package for handling Auth Middleware using JWT
**
**
**
**
 */

func JwtVerifyToken(ctx *fiber.Ctx) error {
	JwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if JwtToken == "" {
		res := response.ResponseDTO{
			Code:    401,
			Message: "Unauthorized",
		}
		return ctx.Status(401).JSON(res)
	}

	req := new(http.Request)
	req.Header = http.Header{}
	req.Header.Set("Authorization", JwtToken)

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(token.Claims.(jwt.MapClaims)["token_type"])
		tokenType := token.Claims.(jwt.MapClaims)["token_type"]

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else if tokenType != "access_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		} else {
			return config.Get().PublicKey, nil
		}
	})

	if err != nil || !token.Valid {
		res := response.ResponseDTO{
			Code:    401,
			Message: err.Error(),
		}
		return ctx.Status(401).JSON(res)
	}

	return ctx.Next()
}

func JwtVerifyRefresh(ctx *fiber.Ctx) error {
	JwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if JwtToken == "" {
		res := response.ResponseDTO{
			Code:    401,
			Message: "Unauthorized",
		}
		return ctx.Status(401).JSON(res)
	}
	req := new(http.Request)
	req.Header = http.Header{}
	req.Header.Set("Authorization", JwtToken)

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		tokenType := token.Claims.(jwt.MapClaims)["token_type"]

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else if tokenType != "refresh_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		} else {
			return config.Get().PublicKey, nil
		}
	})

	if err != nil || !token.Valid {
		res := response.ResponseDTO{
			Code:    401,
			Message: err.Error(),
		}
		return ctx.Status(401).JSON(res)
	}

	// check is refresh_token available in scribleDB?
	userID := token.Claims.(jwt.MapClaims)["id"].(string)

	if userID == "" {
		res := response.ResponseDTO{
			Code:    401,
			Message: "Token not found",
		}
		return ctx.Status(401).JSON(res)
	}
	scribleDB := db.NewScribleClient()
	refreshToken := new(auth.RefreshDTO)
	err = scribleDB.Query().Read("refresh_token", userID, &refreshToken)

	if err != nil || refreshToken.Expired < time.Now().Unix() {
		res := response.ResponseDTO{
			Code:    401,
			Message: "Refresh Token was expired",
		}
		return ctx.Status(401).JSON(res)
	}

	ctx.Context().Request.Header.Set("userID", userID)
	return ctx.Next()
}
