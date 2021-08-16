package main

import (
	"fmt"
	"golang-starter/config"
	"golang-starter/infrastructures/db"
	"golang-starter/infrastructures/local_db"
	"golang-starter/internal/routes"
	"golang-starter/internal/web"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// you didn't define port in env file
	// the default port is random from fiber

	appPort := config.Get().AppPort
	log.Println("Server running on PORT", appPort)
	app := fiber.New()

	mysqlDB := db.NewMysqlClient()
	scribleDB := local_db.NewScribleClient()

	routeStruct := routes.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:       app,
			MysqlDB:   mysqlDB,
			ScribleDB: scribleDB,
		},
	}
	router := routes.NewHttpRoute(routeStruct)
	router.GetRoutes()

	app.Listen(fmt.Sprintf(":%s", appPort))
}
