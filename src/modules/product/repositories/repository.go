package repositories

import (
	"golang-starter/infrastructures/db"

	"github.com/jmoiron/sqlx"
)

type Repositories interface {
	RepositoryProductsCommand
	RepositoryProductsQuery
	RepositoryProductsImagesCommand
	RepositoryProductsImagesQuery
	StartTx() *sqlx.Tx
	CloseTx()
}

type RepositoriesImpl struct {
	// inject db impl to RepositoriesImpl event the db is being used by the child struct impl
	db *db.MysqlImpl
	*RepositoryProductsCommandImpl
	*RepositoryProductsQueryImpl
	*RepositoryProductsImagesCommandImpl
	*RepositoryProductsImagesQueryImpl
}

func NewRepository(
	db *db.MysqlImpl,
) *RepositoriesImpl {
	return &RepositoriesImpl{
		db: db,
		RepositoryProductsQueryImpl: &RepositoryProductsQueryImpl{
			db: db.DB,
		},
		RepositoryProductsCommandImpl: &RepositoryProductsCommandImpl{
			db: db.DB,
		},
		RepositoryProductsImagesCommandImpl: &RepositoryProductsImagesCommandImpl{
			db: db.DB,
		},
		RepositoryProductsImagesQueryImpl: &RepositoryProductsImagesQueryImpl{
			db: db.DB,
		},
	}
}

// StartTx: Inject transaction
// once you defined your repository, you should define your StartTx to using transaction
// so every you create your repository, you should define your StartTx function to handle transactional database
func (repo *RepositoriesImpl) StartTx() *sqlx.Tx {
	tx := repo.db.DB.MustBegin()

	// inject all tx field from struct impl
	repo.RepositoryProductsCommandImpl.tx = tx
	repo.RepositoryProductsImagesCommandImpl.tx = tx

	return tx
}

// CloseTx: close your transaction statement
// once you StartTx, all of your action will be use a transaction.
// and because the transaction has been committed or rollback you will get this error
/// sql: transaction has already been committed or rolled back.
// so you must define your commit and rollback even you don't want to use transaction
func (repo *RepositoriesImpl) CloseTx() {
	// set tx to nil
	repo.RepositoryProductsCommandImpl.tx = nil
	repo.RepositoryProductsImagesCommandImpl.tx = nil
}
