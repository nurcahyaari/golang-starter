package config

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

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
	// Redis
	RedisHost     string
	RedisPort     string
	RedisDB       int
	RedisPassword string
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
			log.Println("Err: Cannot load .env file")
		}

		config = load()
	})
}

func load() appConfigStruct {

	jwtTokenExp := os.Getenv("JWT_TOKEN_EXPIRED")
	jwtRefreshExp := os.Getenv("JWT_REFRESH_EXPIRED")

	jwtTokenDuration, _ := time.ParseDuration(jwtTokenExp)
	jwtRefreshDuration, _ := time.ParseDuration(jwtRefreshExp)

	envRedisDB := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(envRedisDB)
	if err != nil {
		log.Println("Err: Cannot parse redisDB into int")
	}
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
		// redis
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisDB:       redisDB,
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		// key
		PrivateKey: os.Getenv("PRIVATE_KEY"),
		PublicKey:  os.Getenv("PUBLIC_KEY"),
		// Jwt Configuration
		JwtTokenType:      "Bearer",
		JwtTokenExpired:   jwtTokenDuration,   // in second
		JwtRefreshExpired: jwtRefreshDuration, // in second
	}
}

func Get() appConfigStruct {
	return config
}
