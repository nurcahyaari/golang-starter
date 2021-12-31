package http

import (
	httpresponse "golang-starter/internal/protocols/http/response"
	"net/http"
)

func (h HttpHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	// userData := new(dto.UserRequestLoginBody)

	// if err := ctx.BodyParser(userData); err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := h.UserService.Login(userData)
	// if err != nil {
	// 	return httpresponse.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	// }

	// return httpresponse.JsonResponse(ctx, http.StatusOK, "", res)
	httpresponse.Text(w, http.StatusOK, "Success")
}

func (h HttpHandlerImpl) Refresh(w http.ResponseWriter, r *http.Request) {
	// userID := ctx.Get("userID")

	// res, err := h.UserService.RefreshToken(userID)

	// if err != nil {
	// 	return httpresponse.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	// }

	// return httpresponse.JsonResponse(ctx, http.StatusOK, "", res)
	httpresponse.Text(w, http.StatusOK, "Success")
}
