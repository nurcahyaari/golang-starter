package middleware

import (
	"fmt"
	"golang-starter/infrastructure/config"
	"golang-starter/infrastructure/db/localdb"
	"golang-starter/infrastructure/utils/auth"
	"golang-starter/infrastructure/utils/response"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/gofiber/fiber"
)

/*
** Package for handling Auth Middleware using JWT
**
**
**
**
 */

func JwtMiddleware(ctx *fiber.Ctx) {
	fmt.Println("Middleware")
	ctx.Next()
}

func JwtVerifyToken(ctx *fiber.Ctx) {
	JwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if JwtToken == "" {
		res := response.ResponseDTO{
			Code:    401,
			Message: "Unauthorized",
		}
		ctx.Status(401)
		ctx.JSON(res)
		return
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
		ctx.Status(401)
		ctx.JSON(res)
		return
	}

	ctx.Next()
}

func JwtVerifyRefresh(ctx *fiber.Ctx) {
	JwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if JwtToken == "" {
		res := response.ResponseDTO{
			Code:    401,
			Message: "Unauthorized",
		}
		ctx.Status(401)
		ctx.JSON(res)
		return
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
		ctx.Status(401)
		ctx.JSON(res)
		return
	}

	// check is refresh_token available in localDB?
	userID := token.Claims.(jwt.MapClaims)["id"].(string)

	if userID == "" {
		res := response.ResponseDTO{
			Code:    401,
			Message: "Token not found",
		}
		ctx.Status(401)
		ctx.JSON(res)
		return
	}
	localDB := localdb.Load()
	refreshToken := new(auth.RefreshDTO)
	err = localDB.Query().Read("refresh_token", userID, &refreshToken)

	if err != nil || refreshToken.Expired < time.Now().Unix() {
		res := response.ResponseDTO{
			Code:    401,
			Message: "Refresh Token was expired",
		}
		ctx.Status(401)
		ctx.JSON(res)
		return
	}

	ctx.Fasthttp.Request.Header.Set("userID", userID)
	ctx.Next()
}
