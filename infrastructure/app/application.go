package app

import (
	"golang-starter/infrastructure/config"
	"golang-starter/infrastructure/routes"
	"os"

	"github.com/gofiber/fiber"
)

// MainApplication is using for wrapping all applications layer
func MainApplication() error {
	// you didn't define port in env file
	// the default port is random from fiber
	err := config.AppConfig()

	if err != nil {
		return err
	}

	appPort := os.Getenv("APP_PORT")
	// appName := os.Getenv("APP_NAME")
	// fmt.Println(appPort)
	app := fiber.New()

	routes.RegisterRoute(app)

	app.Listen(appPort)

	return nil
}
