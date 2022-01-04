package db

import (
	"fmt"
	"golang-starter/config"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type MysqlImpl struct {
	DB *sqlx.DB
}

func NewMysqlClient() *MysqlImpl {
	log.Info().Msg("Initialize Mysql connection")
	var err error

	dbHost := config.Get().DB.Mysql.Host
	dbPort := config.Get().DB.Mysql.Port
	dbName := config.Get().DB.Mysql.Name
	dbUser := config.Get().DB.Mysql.User
	dbPass := config.Get().DB.Mysql.Pass

	sHost := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("mysql", sHost)

	if err != nil {
		log.Err(err).Msgf("Error to loading Database %s", err)
		panic(err)
	}

	log.Info().Str("Name", dbName).Msg("Success connect to DB")
	return &MysqlImpl{
		DB: db,
	}
}

func (db MysqlImpl) TransactionCallback(transaction func() (interface{}, error)) (interface{}, error) {
	retry := 0
	maxRetry := 1000

	res, err := transaction()

	if err != nil {
		for strings.Contains(err.Error(), "Error 1213") {
			log.Err(err).Msg("Restaring transaction")
			time.Sleep(10 * time.Millisecond)
			res, err = transaction()
			if err != nil {
				fmt.Println(err)
				if !strings.Contains(err.Error(), "Error 1213") {
					if retry >= maxRetry {
						log.Info().Msgf("Retrying transaction %d success", retry)
						break
					}
				} else {
					break
				}
			} else {
				log.Info().Msgf("Retrying transaction %d success", retry)
				break
			}
		}
	}

	return res, err
}
