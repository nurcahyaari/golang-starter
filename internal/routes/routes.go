package routes

import (
	"golang-starter/internal/web"
	productRoute "golang-starter/src/products/router"
	userRoute "golang-starter/src/users/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RouterStruct struct {
	web.RouterStruct
}

func NewHttpRoute(r RouterStruct) RouterStruct {
	log.Println("Loading the HTTP Router")

	return r
}

func (c *RouterStruct) GetRoutes() {
	c.Web.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello this is my first route in go fiber"))
	})

	// registering route from another modules
	productRouterStruct := productRoute.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:       c.Web,
			MysqlDB:   c.MysqlDB,
			ScribleDB: c.ScribleDB,
		},
	}
	productRouter := productRoute.NewHttpRoute(productRouterStruct)
	productRouter.GetRoute()

	userRouterStruct := userRoute.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:       c.Web,
			MysqlDB:   c.MysqlDB,
			ScribleDB: c.ScribleDB,
		},
	}
	userRouter := userRoute.NewHttpRoute(userRouterStruct)
	userRouter.GetRoute()

	// handling 404 error
	c.Web.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
}
