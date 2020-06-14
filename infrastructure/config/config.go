package config

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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
	// database config
	DbDialeg   string
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
	// key
	PrivateKey *rsa.PrivateKey
	// jwt time expired
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
			panic(err)
		}

		config = load()
	})
}

func readPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open("infrastructure/config/public.key") //openssl rsa -in app.rsa -pubout > app.rsa.pub
	if err != nil {
		panic(err)
	}

	pemFileInfo, _ := publicKeyFile.Stat()
	var size = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode([]byte(pemBytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	publicKey, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return publicKey
}

func readPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open("infrastructure/config/private.key")
	if err != nil {
		panic(err)
	}

	pemFileInfo, _ := privateKeyFile.Stat()
	var size = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode([]byte(pemBytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}
	return privateKeyImported
}

func load() appConfigStruct {
	privateKey := readPrivateKey()

	return appConfigStruct{
		AppPort: os.Getenv("APP_PORT"),
		AppKey:  os.Getenv("APP_KEY"),
		// db configure
		DbDialeg:   os.Getenv("DB_DIALEG"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		PrivateKey: privateKey,
		// Jwt Configuration
		JwtTokenExpired: 60, // in second
	}
}

func Get() appConfigStruct {
	return config
}
