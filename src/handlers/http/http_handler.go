package http

import (
	productsvc "golang-starter/src/domains/product/services"
	usersvc "golang-starter/src/domains/user/services"

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
	r.Get("/products/{product_id}", h.GetProductByID)
}
