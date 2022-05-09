//go:build wireinject
// +build wireinject

package main

import (
	"golang-starter/infrastructures/db"
	"golang-starter/infrastructures/db/transaction"
	"golang-starter/infrastructures/localdb"
	"golang-starter/internal/protocols/http"
	httprouter "golang-starter/internal/protocols/http/router"
	jwtauth "golang-starter/internal/utils/auth"
	httphandler "golang-starter/src/handlers/http"
	productrepo "golang-starter/src/modules/product/repositories"
	productsvc "golang-starter/src/modules/product/services"
	userrepo "golang-starter/src/modules/user/repositories"
	usersvc "golang-starter/src/modules/user/services"

	"github.com/google/wire"
)

// wiring transaction
// var transactionDB = wire.NewSet(
// 	transaction.NewTransaction,
// )

// wiring jwt auth
var jwtAuth = wire.NewSet(
	jwtauth.NewJwt,
	wire.Bind(
		new(jwtauth.JwtToken),
		new(*jwtauth.JwtTokenImpl),
	),
)

// Wiring for domain

// product
var productRepo = wire.NewSet(
	productrepo.NewRepository,
	wire.Bind(
		new(productrepo.Repositories),
		new(*productrepo.RepositoriesImpl),
	),
)

var productSvc = wire.NewSet(
	productsvc.NewProductService,
	wire.Bind(
		new(productsvc.ProductService),
		new(*productsvc.ProductServiceImpl),
	),
)

// user
var userMysqlRepo = wire.NewSet(
	userrepo.NewRepository,
	wire.Bind(
		new(userrepo.Repositories),
		new(*userrepo.RepositoriesImpl),
	),
)
var userScribleRepo = wire.NewSet(
	userrepo.NewUserScribleRepository,
	wire.Bind(
		new(userrepo.UserScribleRepository),
		new(*userrepo.UserScribleRepositoryImpl),
	),
)

var userSvc = wire.NewSet(
	usersvc.NewUserService,
	wire.Bind(
		new(usersvc.UserService),
		new(*usersvc.UserServiceImpl),
	),
)

// Wiring for http protocol
var httpHandler = wire.NewSet(
	httphandler.NewHttpHandler,
)

// Wiring protocol routing
var httpRouter = wire.NewSet(
	httprouter.NewHttpRoute,
)

func InitHttpProtocol() *http.HttpImpl {
	wire.Build(
		db.NewMysqlClient,
		localdb.NewScribleClient,
		transaction.NewTransaction,
		productRepo,
		userMysqlRepo,
		userScribleRepo,
		jwtAuth,
		productSvc,
		userSvc,
		httpHandler,
		httpRouter,
		http.NewHttpProtocol,
	)
	return &http.HttpImpl{}
}
