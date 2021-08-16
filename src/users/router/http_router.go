package router

import (
	"golang-starter/internal/middleware"
	"golang-starter/internal/utils/auth"
	"golang-starter/src/users/handlers"
	"golang-starter/src/users/repositories"
	"golang-starter/src/users/services"
	"log"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Users Route")

	structs.jwtAuth = auth.NewJwt(structs.ScribleDB)

	return structs
}

func (r *RouterStruct) GetRoute() {
	userMysqlRepo := repositories.NewUserMysqlRepository(r.MysqlDB)
	userScribleRepo := repositories.NewUserScribleRepositoryInterface(r.ScribleDB)
	userService := services.NewUserService(userMysqlRepo, r.jwtAuth, userScribleRepo)
	userHandlers := handlers.NewHttpHandler(userService)

	r.Web.Post("/user/login", userHandlers.Login)
	r.Web.Post("/user/refresh-token", middleware.JwtVerifyRefresh, userHandlers.Refresh)
}
