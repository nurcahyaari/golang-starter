package router

import (
	"golang-starter/internal/middleware"
	"golang-starter/src/products/handlers"
	"golang-starter/src/products/repositories"
	"golang-starter/src/products/services"
	"log"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Products Route")
	return structs
}

func (r *RouterStruct) GetRoute() {
	productRepository := repositories.NewProductRepostiory(r.MysqlDB)
	productService := services.NewProductService(productRepository)
	productHandlers := handlers.NewHttpHandler(productService)

	r.Web.Get("/products", middleware.JwtVerifyToken, productHandlers.GetProducts)
}
