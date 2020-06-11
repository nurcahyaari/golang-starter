package config

import (
	"log"

	"github.com/pkg/errors"

	"github.com/joho/godotenv"
)

// AppConfig is module for bundle Application Config
func AppConfig() error {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Cannot Load .env file")
		return errors.New("Cannot Load .env file")
	}
	return nil
}
