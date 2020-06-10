package main

import (
	"golang-starter/infrastructure/app"
	"log"
)

func main() {
	log.Print("Server is starting\n\n")
	app.MainApplication()
}
