package routes

import (
	"example/pawpi/constants"
	"example/pawpi/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	e.GET("/users", controller.GetManyController)
	e.GET("/users/:id", controller.GetUserDetailController)
	e.POST("/login", controller.LoginUsersController)

	//Login
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users/:id", controller.GetUserDetailController)
	eJwt.DELETE("/users/:id", controller.DeleteControllerAuth)
	eJwt.PUT("/users/:id", controller.UpdateControllerAuth)
	eJwt.POST("/books", controller.CreateController2)
	eJwt.DELETE("/books/:id", controller.DeleteController2)
	eJwt.PUT("/books/:id", controller.UpdateController2)

	//Biasa
	e.POST("/users", controller.CreateController)
	e.DELETE("/users/:id", controller.DeleteController)
	e.PUT("/users/:id", controller.UpdateController)

	//books
	e.GET("/books", controller.GetManyController2)
	e.GET("/books/:id", controller.GetOneController2)

}
