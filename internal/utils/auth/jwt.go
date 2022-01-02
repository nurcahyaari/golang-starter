package auth

import (
	"fmt"
	"golang-starter/config"
<<<<<<< HEAD
	"golang-starter/infrastructures/localdb"
	"golang-starter/infrastructures/logger"
=======
	localdb "golang-starter/infrastructures/local_db"
>>>>>>> 69dd86c29b7f455fad88cb5217980278d0199b8a

	"golang-starter/internal/utils/auth/dto"
	"golang-starter/internal/utils/encryption"
	"golang-starter/internal/utils/rsa"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

type JwtToken interface {
	Sign(claims jwt.MapClaims) dto.Token
	SignRSA(claims jwt.MapClaims) dto.Token
}

type JwtTokenImpl struct {
	jwtTokenTimeExp        time.Duration
	jwtRefreshTokenTimeExp time.Duration
	cached                 *localdb.ScribleImpl
}

func NewJwt(cached *localdb.ScribleImpl) *JwtTokenImpl {
	jwtTokenDuration, err := time.ParseDuration(config.Get().Auth.JwtToken.Expired)
	if err != nil {
		log.Err(err).Msg(config.Get().Auth.JwtToken.Expired)
	}
	jwtRefreshDuration, err := time.ParseDuration(config.Get().Auth.JwtToken.RefreshExpired)
	if err != nil {
		log.Err(err).Msg(config.Get().Auth.JwtToken.RefreshExpired)
	}
	return &JwtTokenImpl{
		cached:                 cached,
		jwtTokenTimeExp:        jwtTokenDuration,
		jwtRefreshTokenTimeExp: jwtRefreshDuration,
	}
}

func (o JwtTokenImpl) Sign(claims jwt.MapClaims) dto.Token {
	timeNow := time.Now()
	tokenExpired := timeNow.Add(o.jwtTokenTimeExp).Unix()

	if claims["id"] == nil {
		return dto.Token{}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// setup userdata
	var _, checkExp = claims["exp"]
	var _, checkIat = claims["exp"]

	// if user didn't define claims expired
	if !checkExp {
		claims["exp"] = tokenExpired
	}
	// if user didn't define claims iat
	if !checkIat {
		claims["iat"] = timeNow.Unix()
	}
	claims["token_type"] = "access_token"

	token.Claims = claims

	authToken := new(dto.Token)
	tokenString, err := token.SignedString([]byte(config.Get().Application.Key.Default))

	if err != nil {
		log.Err(err)
		return dto.Token{}
	}

	authToken.Token = tokenString
	authToken.Type = "Bearer"

	//create refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenExpired := timeNow.Add(o.jwtRefreshTokenTimeExp).Unix()

	claims["exp"] = refreshTokenExpired
	claims["token_type"] = "refresh_token"
	refreshToken.Claims = claims

	refreshTokenString, err := refreshToken.SignedString([]byte(config.Get().Application.Key.Default))

	if err != nil {
		return dto.Token{}
	}
	authToken.RefreshToken = refreshTokenString

	//save token to redis db
	go func() {
		encryptedRefreshToken, err := encryption.AesCFBEncryption(refreshTokenString, config.Get().Application.Key.Default)
		if err != nil {
			log.Err(err)
		}
		// check data type of the claims
		switch claims["id"].(type) {
		case int:
			claims["id"] = fmt.Sprintf("%d", claims["id"].(int))
		case int32:
			claims["id"] = fmt.Sprintf("%d", claims["id"].(int32))
		case float64:
			claims["id"] = fmt.Sprintf("%d", int(claims["id"].(float64)))
		default:
		}
		o.cached.DB().Write("refresh_token", claims["id"].(string), dto.RefreshToken{RefreshToken: encryptedRefreshToken, Expired: refreshTokenExpired})
		if err != nil {
			log.Err(err).Msgf("Failed to save refresh token to scrible")
		} else {
			log.Info().Msg("Successfully to save refresh token to scrible")
		}
	}()

	return dto.Token{
		Type:         config.Get().Auth.JwtToken.Type,
		Token:        authToken.Token,
		RefreshToken: authToken.RefreshToken,
	}
}

// Sign ins method to generate jwt token and refresh token
// it has ... parameter
// userdata is map data, it's using for passing user data
// default expired time is 60 second
func (o JwtTokenImpl) SignRSA(claims jwt.MapClaims) dto.Token {
	timeNow := time.Now()
	tokenExpired := timeNow.Add(o.jwtTokenTimeExp).Unix()

	if claims["id"] == nil {
		return dto.Token{}
	}

	token := jwt.New(jwt.SigningMethodRS256)
	// setup userdata
	var _, checkExp = claims["exp"]
	var _, checkIat = claims["exp"]

	// if user didn't define claims expired
	if !checkExp {
		claims["exp"] = tokenExpired
	}
	// if user didn't define claims iat
	if !checkIat {
		claims["iat"] = timeNow.Unix()
	}
	claims["token_type"] = "access_token"

	token.Claims = claims
	authToken := new(dto.Token)
	privateRsa, err := rsa.ReadPrivateKeyFromEnv(config.Get().Application.Key.Rsa.Private)
	if err != nil {
		log.Err(err).Msg("err read private key rsa from env")
		return dto.Token{}
	}
	tokenString, err := token.SignedString(privateRsa)
	if err != nil {
		log.Err(err).Msg("err read private rsa")
		return dto.Token{}
	}

	authToken.Token = tokenString
	authToken.Type = "Bearer"

	//create refresh token
	refreshToken := jwt.New(jwt.SigningMethodRS256)
	refreshTokenExpired := timeNow.Add(o.jwtRefreshTokenTimeExp).Unix()

	claims["exp"] = refreshTokenExpired
	claims["token_type"] = "refresh_token"
	refreshToken.Claims = claims
	refreshTokenString, err := refreshToken.SignedString(privateRsa)
	if err != nil {
		log.Err(err).Msg("")
		return dto.Token{}
	}
	authToken.RefreshToken = refreshTokenString

	//save token to redis db
	go func() {
		encryptedRefreshToken, err := encryption.AesCFBEncryption(refreshTokenString, config.Get().Application.Key.Default)
		if err != nil {
			log.Err(err)
		}
		// check data type of the claims
		switch claims["id"].(type) {
		case int:
			claims["id"] = fmt.Sprintf("%d", claims["id"].(int))
		case int32:
			claims["id"] = fmt.Sprintf("%d", claims["id"].(int32))
		case float64:
			claims["id"] = fmt.Sprintf("%d", int(claims["id"].(float64)))
		default:
		}
		o.cached.DB().Write("refresh_token", claims["id"].(string), dto.RefreshToken{RefreshToken: encryptedRefreshToken, Expired: refreshTokenExpired})
		if err != nil {
			log.Err(err).Msg("Failed to save refresh token to redis")
		} else {
			log.Info().Msg("Successfully to save refresh token to redis")
		}
	}()

	return dto.Token{
		Type:         "Bearer",
		Token:        authToken.Token,
		RefreshToken: authToken.RefreshToken,
	}
}
