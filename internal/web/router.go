package web

import (
	"golang-starter/infrastructures/db"
	"golang-starter/infrastructures/local_db"

	"github.com/gofiber/fiber/v2"
)

type RouterStruct struct {
	Web       *fiber.App
	MysqlDB   db.MysqlDB
	ScribleDB local_db.ScribleDB
}
