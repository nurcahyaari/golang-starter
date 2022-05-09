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
		s.Rollback()
		return err
	}

	return s.Commit()
}
