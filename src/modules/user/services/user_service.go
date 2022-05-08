package services

import (
	"context"
	"golang-starter/internal/protocols/http/errors"
	"golang-starter/internal/utils/auth"
	"golang-starter/src/modules/user/dto"
	"golang-starter/src/modules/user/repositories"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindByID(ctx context.Context, id uint) (*dto.UserRespBody, error)
	UserLogin(ctx context.Context, req dto.UserRequestLoginBody) (*dto.UserTokenRespBody, error)
	UserRefreshToken(ctx context.Context, userId string) (*dto.UserTokenRespBody, error)
}

type UserServiceImpl struct {
	userRepository repositories.Repositories
	jwtAuth        auth.JwtToken
}

func NewUserService(
	jwtAuth auth.JwtToken,
	userRepository repositories.Repositories,
	userScribleRepository repositories.UserScribleRepository,
) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
		jwtAuth:        jwtAuth,
	}
}

func (s UserServiceImpl) FindByID(ctx context.Context, userId uint) (*dto.UserRespBody, error) {
	user, err := s.userRepository.
		FilterUsers(repositories.NewUsersFilter("AND").SetFilterByUserId(userId, "=")).
		GetUsers(ctx)

	if err != nil {
		log.Err(err).Msg("error fetch user data")
		return nil, errors.FindErrorType(err)
	}

	userResp := dto.CreateUserResp(*user)

	return &userResp, nil
}

func (s UserServiceImpl) UserLogin(ctx context.Context, req dto.UserRequestLoginBody) (*dto.UserTokenRespBody, error) {

	user, err := s.userRepository.
		FilterUsers(repositories.NewUsersFilter("AND").SetFilterByEmail(req.Email, "=")).
		GetUsers(ctx)
	if err != nil {
		log.Err(err).Msg("error fetch user data")
		return nil, errors.FindErrorType(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.Unauthorization("email and password didn't match")
	}

	userToken := s.jwtAuth.SignRSA(jwt.MapClaims{
		"id": user.UserId,
	})

	token := dto.UserTokenRespBody(userToken)

	return &token, nil
}

func (s UserServiceImpl) UserRefreshToken(ctx context.Context, userId string) (*dto.UserTokenRespBody, error) {
	refreshToken, err := s.userRepository.FindUserRefreshToken(userId)
	if err != nil {
		return nil, errors.Unauthorization("token is not valid")
	}

	if refreshToken.Expired < time.Now().Unix() {
		return nil, errors.Unauthorization("token is expired")
	}

	userToken := s.jwtAuth.SignRSA(jwt.MapClaims{
		"id": userId,
	})

	token := dto.UserTokenRespBody(userToken)

	return &token, nil
}
