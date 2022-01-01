package http

import (
	"encoding/json"
	"fmt"
	httpresponse "golang-starter/internal/protocols/http/response"
	"golang-starter/src/modules/user/dto"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (h HttpHandlerImpl) GetdUserById(w http.ResponseWriter, r *http.Request) {
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

func (h HttpHandlerImpl) UserRefreshToken(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("id")
	fmt.Println(userId)

	res, err := h.UserService.UserRefreshToken(r.Context(), userId)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", res)
}
