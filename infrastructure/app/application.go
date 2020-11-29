package app

import (
	"fmt"
	"golang-starter/infrastructure/config"
	"golang-starter/infrastructure/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

// MainApplication is using for wrapping all applications layer
func MainApplication() {
	// you didn't define port in env file
	// the default port is random from fiber

	appPort := config.Get().AppPort
	log.Println("Server running on PORT", appPort)
	app := fiber.New()

	routes.RegisterRoute(app)

	app.Listen(fmt.Sprintf(":%s", appPort))
}
