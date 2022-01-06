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
}

type RepositoriesImpl struct {
	*RepositoryProductsCommandImpl
	*RepositoryProductsQueryImpl
	*RepositoryProductsImagesCommandImpl
	*RepositoryProductsImagesQueryImpl
}

func NewRepository(
	db *db.MysqlImpl,
) *RepositoriesImpl {
	return &RepositoriesImpl{
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
	tx := db.NewMysqlTx().Tx

	// inject all tx field from struct impl
	repo.RepositoryProductsCommandImpl = &RepositoryProductsCommandImpl{
		tx: tx,
	}
	repo.RepositoryProductsImagesCommandImpl = &RepositoryProductsImagesCommandImpl{
		tx: tx,
	}
	return tx
}
