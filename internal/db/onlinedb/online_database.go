package onlinedb

import (
	"fmt"
	"golang-starter/infrastructure/config"
	"log"

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
	// log.Println("Initialize Database connection")
	var err error
	var db *gorm.DB
	dbDialeg := config.Get().DbDialeg
	dbHost := config.Get().DbHost
	dbPort := config.Get().DbPort
	dbName := config.Get().DbName
	dbUser := config.Get().DbUsername
	dbPassword := config.Get().DbPassword

	sHost := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	sDialeg := fmt.Sprintf("%s", dbDialeg)

	db, err = gorm.Open(sDialeg, sHost)
	// defer db.Close()

	if err != nil {
		log.Fatal("Error to loading Database %s", err)
		panic(err)
	}
	// log.Println("Database was connected")
	return &database{
		db: db,
	}
}

func (database *database) Query() *gorm.DB {
	return database.db
}
