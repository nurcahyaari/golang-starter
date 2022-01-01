package middleware

import (
	"fmt"
	"golang-starter/config"
	"net/http"
	"strings"
	"time"

	httpresponse "golang-starter/internal/protocols/http/response"
	"golang-starter/internal/utils/rsa"

	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"github.com/rs/zerolog/log"
)

/*
** Package for handling Auth Middleware using JWT
**
**
**
**
 */
// JwtVerifyToken usefull for middleware for verify the jwt token from the Authorization
// this function will serve to middleware and usefull for the idiomatic framework like gorm or chi or just net/http
func JwtVerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JwtToken := strings.Replace(r.Header.Get("Authorization"), fmt.Sprintf("%s ", "Bearer"), "", 1)

		if JwtToken == "" {
			httpresponse.Json(w, http.StatusUnauthorized, "", "token is empty")
			return
		}

		token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			tokenType := token.Claims.(jwt.MapClaims)["token_type"]

			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			} else if tokenType != "access_token" {
				return nil, fmt.Errorf("unexpected token type: %v", tokenType)
			} else {
				publicRsa, err := rsa.ReadPublicKeyFromEnv(config.Get().Application.Key.Rsa.Public)
				if err != nil {
					log.Err(err).Msg("err read private key rsa from env")
					return nil, nil
				}
				return publicRsa, nil
			}
		})

		if err != nil || !token.Valid {
			log.Err(err)
			httpresponse.Json(w, http.StatusUnauthorized, "", "Token is not valid")
			return
		}

		rawId := token.Claims.(jwt.MapClaims)["id"].(float64)
		id := fmt.Sprintf("%d", int(rawId))
		if id == "" {
			httpresponse.Json(w, http.StatusUnauthorized, "", "Token not Found")
			return
		}

		rawExp := token.Claims.(jwt.MapClaims)["exp"].(float64)
		exp := int64(rawExp)
		if exp < time.Now().Unix() {
			httpresponse.Json(w, http.StatusUnauthorized, "", "Token has expired")
			return
		}

		r.Header.Set("id", id)

		next.ServeHTTP(w, r)
	})
}

// JwtVerifyRefreshToken usefull for middleware for verify the jwt refresh token from the Authorization
// this function will serve to middleware and usefull for the idiomatic framework like gorm or chi or just net/http
func JwtVerifyRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JwtToken := strings.Replace(r.Header.Get("Authorization"), fmt.Sprintf("%s ", "Bearer"), "", 1)

		if JwtToken == "" {
			httpresponse.Json(w, http.StatusUnauthorized, "", "Refresh token is empty")
			return
		}

		token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			tokenType := token.Claims.(jwt.MapClaims)["token_type"]
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			} else if tokenType != "refresh_token" {
				return nil, fmt.Errorf("unexpected token type: %v", tokenType)
			} else {
				privateRsa, err := rsa.ReadPublicKeyFromEnv(config.Get().Application.Key.Rsa.Public)
				if err != nil {
					log.Err(err).Msg("err read private key rsa from env")
					return nil, nil
				}
				return privateRsa, nil
			}
		})

		if err != nil || !token.Valid {
			log.Err(err).Msg("")
			httpresponse.Json(w, http.StatusUnauthorized, "", "Refresh token is not valid")
			return
		}

		rawId := token.Claims.(jwt.MapClaims)["id"].(float64)
		id := fmt.Sprintf("%d", int(rawId))
		if id == "" {
			httpresponse.Json(w, http.StatusUnauthorized, "", "Refresh token not Found")
			return
		}

		rawExp := token.Claims.(jwt.MapClaims)["exp"].(float64)
		exp := int64(rawExp)
		if exp < time.Now().Unix() {
			httpresponse.Json(w, http.StatusUnauthorized, "", "Refresh token has expired")
			return
		}

		r.Header.Set("id", id)

		next.ServeHTTP(w, r)
	})
}
