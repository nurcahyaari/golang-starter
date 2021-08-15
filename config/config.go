package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/joho/godotenv"
)

var config appConfigStruct
var doOnce sync.Once

type appConfigStruct struct {
	AppPort string
	AppKey  string // all off local encryption will use this key
	LogPath string
	// database config
	DbDialeg   string
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
	// key
	PrivateKey string
	PublicKey  string
	// jwt
	JwtTokenType      string
	JwtTokenExpired   time.Duration // in second
	JwtRefreshExpired time.Duration // in second
}

func init() {
	doOnce.Do(func() {
		err := godotenv.Load()
		log.Println("Loading .env file....")
		if err != nil {
			log.Fatalln("Cannot Load .env file")
			errors.New("Cannot Load .env file")
		}

		config = load()
	})
}

func load() appConfigStruct {
	return appConfigStruct{
		AppPort: os.Getenv("APP_PORT"),
		AppKey:  os.Getenv("APP_KEY"),
		LogPath: os.Getenv("LOG_PATH"),
		// db configure
		DbDialeg:   os.Getenv("DB_DIALEG"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		PrivateKey: os.Getenv("PRIVATE_KEY"),
		PublicKey:  os.Getenv("PUBLIC_KEY"),
		// Jwt Configuration
		JwtTokenType:      "Bearer",
		JwtTokenExpired:   60 * 60,           // in second
		JwtRefreshExpired: 60 * 60 * 24 * 30, // in second
	}
}

func Get() appConfigStruct {
	return config
}
