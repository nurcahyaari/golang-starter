package router

import (
	"golang-starter/internal/db"
	"golang-starter/internal/middleware"
	"golang-starter/src/users/controllers"
	"golang-starter/src/users/repositories"
	"golang-starter/src/users/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	db := db.NewMysqlClient()

	userRepo := repositories.ProvideUserRepository(db)
	userService := services.ProvideUserService(userRepo)
	userController := controllers.ProvideUserController(userService)

	app.Post("/user/login", userController.Login)
	app.Post("/user/refresh-token", middleware.JwtVerifyRefresh, userController.Refresh)
}
