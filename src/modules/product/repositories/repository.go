package repositories

import "golang-starter/infrastructures/db"

type Repositories interface {
	RepositoryProductsCommand
	RepositoryProductsQuery
}

type RepositoriesImpl struct {
	*RepositoryProductsCommandImpl
	*RepositoryProductsQueryImpl
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
	}
}
