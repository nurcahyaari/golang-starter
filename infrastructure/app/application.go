package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

// MainApplication is using for wrapping all applications layer
func MainApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// you didn't define port in env file
	// the default port is random from fiber
	appPort := os.Getenv("APP_PORT")
	// appName := os.Getenv("APP_NAME")

	app := fiber.New()

	// routes.RegisterRoute(app)

	app.Listen(appPort)
}
