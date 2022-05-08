package transaction

import (
	"context"
	"golang-starter/infrastructures/db"
)

type TransactionImpl struct {
	db *db.MysqlImpl
}

func NewTransaction(db *db.MysqlImpl) *TransactionImpl {
	tx := &TransactionImpl{db: db}
	return tx
}

func (s *TransactionImpl) StartTx() {
	err := s.db.DB.Beginx()
	if err != nil {
		return
	}
}

func (s *TransactionImpl) Commit() error {
	return s.db.DB.Commit()
}

func (s *TransactionImpl) Rollback() error {
	return s.db.DB.Rollback()
}

func (s *TransactionImpl) RunWithTransaction(ctx context.Context, f func() error) error {
	s.StartTx()

	err := f()
	if err != nil {
		s.Commit()
		return err
	}

	return s.Rollback()
}

// import (
// 	"fmt"
// 	"strings"
// 	"time"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/rs/zerolog/log"
// )

// type Transaction interface {
// 	TransactionStart() error
// 	TransactionCommit() error
// 	TransactionRollback() error
// 	TransactionCallback(transaction func() (interface{}, error)) (interface{}, error)
// }

// type TransactionImpl struct {
// 	Db *sqlx.DB
// 	Tx *sqlx.Tx
// }

// func NewTransaction(db *sqlx.DB) *TransactionImpl {
// 	tx := db.MustBegin()
// 	return &TransactionImpl{
// 		Tx: tx,
// 	}
// }

// // TransactionCallback is transaction group and using callback function to run the transaction
// // you can use this function same as when you use the transaction without callback. so, if you want to use
// // TransactionCallback just wrap your function with this method
// func (txi TransactionImpl) TransactionCallback(transaction func() (interface{}, error)) (interface{}, error) {
// 	retry := 0
// 	maxRetry := 1000

// 	res, err := transaction()

// 	if err != nil {
// 		for strings.Contains(err.Error(), "Error 1213") {
// 			log.Err(err).Msg("Restaring transaction")
// 			time.Sleep(10 * time.Millisecond)
// 			res, err = transaction()
// 			if err != nil {
// 				fmt.Println(err)
// 				if !strings.Contains(err.Error(), "Error 1213") {
// 					if retry >= maxRetry {
// 						log.Info().Msgf("Retrying transaction %d success", retry)
// 						break
// 					}
// 				} else {
// 					break
// 				}
// 			} else {
// 				log.Info().Msgf("Retrying transaction %d success", retry)
// 				break
// 			}
// 		}
// 	}

// 	return res, err
// }
