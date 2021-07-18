package db

import (
	"fmt"
	"golang-starter/internal/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDB interface {
	Query() *gorm.DB
}

type mysqlDB struct {
	db *gorm.DB
}

func NewMysqlClient() MysqlDB {
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
		return nil
	}
	// log.Println("Database was connected")
	return &mysqlDB{
		db: db,
	}
}

func (mysql mysqlDB) Query() *gorm.DB {
	return mysql.db
}
