package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbDialeg := os.Getenv("DB_DIALEG")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	sHost := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	sDialeg := fmt.Sprintf("%s", dbDialeg)
	fmt.Println(sHost)
	fmt.Println(sDialeg)
	db, err = gorm.Open(sDialeg, sHost)
	// defer db.Close()

	if err != nil {
		log.Fatal("Error to loading Database %s", err)
	}
	log.Println("Database was connected")
	// Database.db = gorm.Open(sDialeg, sHost)

}

func Query() *gorm.DB {
	return db
}
