package main

import (
	"golang-starter/infrastructure/app"
	"log"
)

func main() {
	log.Println("Server is starting")
	app.MainApplication()
}
