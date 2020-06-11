package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database interface {
	Query() *gorm.DB
}

type database struct {
	db *gorm.DB
}

// var db *gorm.DB

func Load() Database {
	log.Println("Initialize Database connection")
	var err error
	var db *gorm.DB
	dbDialeg := os.Getenv("DB_DIALEG")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	sHost := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	sDialeg := fmt.Sprintf("%s", dbDialeg)

	db, err = gorm.Open(sDialeg, sHost)
	// defer db.Close()

	if err != nil {
		log.Fatal("Error to loading Database %s", err)
	}
	log.Println("Database was connected")
	return &database{
		db: db,
	}
}

func (database *database) Query() *gorm.DB {
	return database.db
}

// func Query() *gorm.DB {
// 	return db
// }
