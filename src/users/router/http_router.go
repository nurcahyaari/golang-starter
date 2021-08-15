package router

import (
	"golang-starter/internal/middleware"
	"golang-starter/src/users/handlers"
	"golang-starter/src/users/repositories"
	"golang-starter/src/users/services"
	"log"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Users Route")
	return structs
}

func (r *RouterStruct) GetRoute() {
	userRepo := repositories.NewUserRepository(r.MysqlDB)
	userService := services.NewUserService(userRepo, r.jwtAuth)
	userHandlers := handlers.NewHttpHandler(userService)

	r.Web.Post("/user/login", userHandlers.Login)
	r.Web.Post("/user/refresh-token", middleware.JwtVerifyRefresh, userHandlers.Refresh)
}
