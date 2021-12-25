package db

import (
	"fmt"
	"golang-starter/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlImpl struct {
	db *gorm.DB
}

func NewMysqlClient() *MysqlImpl {
	// log.Println("Initialize Database connection")
	var err error
	var db *gorm.DB
	dbDialeg := "mysql"
	dbHost := config.Get().DB.Mysql.Host
	dbPort := config.Get().DB.Mysql.Port
	dbName := config.Get().DB.Mysql.Name
	dbUser := config.Get().DB.Mysql.User
	dbPassword := config.Get().DB.Mysql.Pass

	sHost := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	sDialeg := dbDialeg

	db, err = gorm.Open(sDialeg, sHost)
	// defer db.Close()

	if err != nil {
		log.Println(fmt.Sprintf("Error to loading Database %s", err))
		panic(err)
	}
	// log.Println("Database was connected")
	return &MysqlImpl{
		db: db,
	}
}

func (c MysqlImpl) DB() *gorm.DB {
	return c.db
}

// // TransactionGroup is a transaction group method to grouping a transaction
// // this function provides the deadlock prevent with max retry 1000
// // this function return an interface and an error
// func (c mysqlDB) TransactionGroup(transaction func() (interface{}, error)) (interface{}, error) {
// 	retry := 0
// 	maxRetry := 1000

// 	res, err := transaction()

// 	if err != nil {
// 		logger.Log.Error("Error in Transaction")
// 		logger.Log.Error(err.Error())
// 		for strings.Contains(err.Error(), "Error 1213") {
// 			logger.Log.Infoln("Restaring transaction")
// 			time.Sleep(10 * time.Millisecond)
// 			res, err = transaction()
// 			if err != nil {
// 				fmt.Println(err)
// 				if !strings.Contains(err.Error(), "Error 1213") {
// 					if retry >= maxRetry {
// 						logger.Log.Info(fmt.Sprintf("Retrying transaction %d success", retry))
// 						break
// 					}
// 				} else {
// 					break
// 				}
// 			} else {
// 				logger.Log.Info(fmt.Sprintf("Retrying transaction %d success", retry))
// 				break
// 			}
// 		}
// 	}

// 	return res, err
// }
