package web

import (
	"golang-starter/infrastructures/db"
	"golang-starter/infrastructures/localdb"

	"github.com/gofiber/fiber/v2"
)

type RouterStruct struct {
	Web       *fiber.App
	MysqlDB   db.MysqlDB
	ScribleDB localdb.ScribleDB
}
