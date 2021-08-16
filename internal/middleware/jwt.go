package middleware

import (
	"fmt"
	"strconv"

	"golang-starter/config"

	"golang-starter/internal/web"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"

	"github.com/gofiber/fiber/v2"
)

/*
** Package for handling Auth Middleware using JWT
**
**
**
**
 */
func JwtVerifyTokenRSA(ctx *fiber.Ctx) error {
	JwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if JwtToken == "" {
		res := web.Response{
			Code:    401,
			Message: "Unauthorized",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	req := new(http.Request)
	req.Header = http.Header{}
	req.Header.Set("Authorization", JwtToken)

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
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
		res := web.Response{
			Code:    401,
			Message: err.Error(),
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	return ctx.Next()
}

func JwtVerifyRefreshRSA(ctx *fiber.Ctx) error {
	JwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if JwtToken == "" {
		res := web.Response{
			Code:    401,
			Message: "Unauthorized",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
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
		res := web.Response{
			Code:    401,
			Message: err.Error(),
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	userID := token.Claims.(jwt.MapClaims)["id"].(string)

	if userID == "" {
		res := web.Response{
			Code:    401,
			Message: "Token not found",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}
	rawExp := token.Claims.(jwt.MapClaims)["exp"].(float64)
	exp := int64(rawExp)

	if exp < time.Now().Unix() {
		res := web.Response{
			Code:    401,
			Message: "Refresh Token was expired",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	ctx.Context().Request.Header.Set("userID", userID)
	return ctx.Next()
}

func JwtVerifyToken(ctx *fiber.Ctx) error {
	jwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if jwtToken == "" {
		res := web.Response{
			Code:    401,
			Message: "Unauthorized",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	req := new(http.Request)
	req.Header = http.Header{}
	req.Header.Set("Authorization", jwtToken)

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]

		if tokenType != "access_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		}
		return []byte(config.Get().AppKey), nil
	})

	if err != nil || !token.Valid {
		res := web.Response{
			Code:    401,
			Message: err.Error(),
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	return ctx.Next()
}

func JwtVerifyRefresh(ctx *fiber.Ctx) error {
	jwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	if jwtToken == "" {
		res := web.Response{
			Code:    401,
			Message: "Unauthorized",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	req := new(http.Request)
	req.Header = http.Header{}
	req.Header.Set("Authorization", jwtToken)

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]
		fmt.Println(tokenType != "refresh_token")
		if tokenType != "refresh_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		}
		return []byte(config.Get().AppKey), nil
	})

	if err != nil || !token.Valid {
		res := web.Response{
			Code:    401,
			Message: err.Error(),
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	rawUserID := token.Claims.(jwt.MapClaims)["id"].(float64)

	if rawUserID == 0 {
		res := web.Response{
			Code:    401,
			Message: "Token not found",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}
	rawExp := token.Claims.(jwt.MapClaims)["exp"].(float64)
	exp := int64(rawExp)

	if exp < time.Now().Unix() {
		res := web.Response{
			Code:    401,
			Message: "Refresh Token was expired",
		}
		return web.JsonResponse(ctx, res.Code, res.Message, nil)
	}

	userID := strconv.Itoa(int(rawUserID))
	ctx.Context().Request.Header.Set("userID", userID)

	return ctx.Next()
}
