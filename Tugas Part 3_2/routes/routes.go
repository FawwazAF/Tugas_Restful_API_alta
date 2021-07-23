package routes

import (
	"example/pawpi/controller"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.GET("/users", controller.GetManyController)
	e.GET("/users/:id", controller.GetOneController)
	e.POST("/users", controller.CreateController)
	e.DELETE("/users/:id", controller.DeleteController)
	e.PUT("/users/:id", controller.UpdateController)

	e.GET("/books", controller.GetManyController2)
	e.GET("/books/:id", controller.GetOneController2)
	e.POST("/books", controller.CreateController2)
	e.DELETE("/books/:id", controller.DeleteController2)
	e.PUT("/books/:id", controller.UpdateController2)

}
