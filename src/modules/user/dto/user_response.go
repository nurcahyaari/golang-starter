package dto

import "golang-starter/src/modules/user/entities"

type UserRespBody struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func CreateUserResp(user entities.Users) UserRespBody {
	return UserRespBody{
		UserID: int(user.UserId),
		Name:   user.Name,
		Email:  user.Email,
	}
}

type UserTokenRespBody struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
