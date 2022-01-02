package http

import (
	"golang-starter/internal/protocols/http/middleware"
	productsvc "golang-starter/src/modules/product/services"
	usersvc "golang-starter/src/modules/user/services"

	"github.com/go-chi/chi/v5"
)

type HttpHandlerImpl struct {
	productsvc.ProductService
	usersvc.UserService
}

func NewHttpHandler(
	productService productsvc.ProductService,
	userService usersvc.UserService,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		ProductService: productService,
		UserService:    userService,
	}
}

func (h *HttpHandlerImpl) Router(r *chi.Mux) {
	r.Get("/products", h.GetProducts)
	r.Get("/products/{productId}", h.GetProductByID)
	r.Get("/users/{userId}", h.GetUserById)
	r.Post("/users/login", h.UserLogin)
	r.With(middleware.JwtVerifyRefreshToken).Post("/users/refresh", h.UserRefreshToken)
}
