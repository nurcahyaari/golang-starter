package repositories

import (
	"golang-starter/infrastructures/db"
)

type Repositories interface {
	RepositoryProductsCommand
	RepositoryProductsQuery
	RepositoryProductsImagesCommand
	RepositoryProductsImagesQuery
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
