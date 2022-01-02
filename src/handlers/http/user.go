package http

import (
	"encoding/json"
	httpresponse "golang-starter/internal/protocols/http/response"
	"golang-starter/src/modules/user/dto"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

// GetUserById return User by userId
// @Summary Get User by userId
// @Description get User by userId
// @Tags Users
// @Param userId path string true "userId"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{userId} [GET]
func (h HttpHandlerImpl) GetUserById(w http.ResponseWriter, r *http.Request) {
	rawUserId := chi.URLParam(r, "userId")
	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	user, err := h.UserService.FindByID(r.Context(), uint(userId))
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", user)
}

// UserLogin login user to get token
// @Summary login user to get token
// @Description login user to get token
// @Tags Users
// @Param User Form body dto.UserRequestLoginBody true "user form"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/login [POST]
func (h HttpHandlerImpl) UserLogin(w http.ResponseWriter, r *http.Request) {
	// userData := new(dto.UserRequestLoginBody)
	userReq := dto.UserRequestLoginBody{}
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	res, err := h.UserService.UserLogin(r.Context(), userReq)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", res)
}

// UserRefreshToken get new token by refresh token
// @Summary get new token by refresh token
// @Description get new token by refresh token
// @Tags Users
// @Param Authorization header string true "refresh token"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/refresh [POST]
func (h HttpHandlerImpl) UserRefreshToken(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("id")

	res, err := h.UserService.UserRefreshToken(r.Context(), userId)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", res)
}
