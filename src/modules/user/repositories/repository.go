package repositories

import (
	"golang-starter/infrastructures/db"
	"golang-starter/infrastructures/localdb"
)

type Repositories interface {
	RepositoryUsersCommand
	RepositoryUsersQuery
	UserScribleRepository
}

type RepositoriesImpl struct {
	*RepositoryUsersCommandImpl
	*RepositoryUsersQueryImpl
	*UserScribleRepositoryImpl
}

func NewRepository(
	db *db.MysqlImpl,
	scribleDB *localdb.ScribleImpl,
) *RepositoriesImpl {
	return &RepositoriesImpl{
		RepositoryUsersCommandImpl: &RepositoryUsersCommandImpl{db: db.DB},
		RepositoryUsersQueryImpl:   &RepositoryUsersQueryImpl{db: db.DB},
		UserScribleRepositoryImpl:  NewUserScribleRepository(scribleDB),
	}
}
