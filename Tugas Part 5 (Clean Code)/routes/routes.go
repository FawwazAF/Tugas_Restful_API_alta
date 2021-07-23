package routes

import (
	"example/pawpi/constants"
	"example/pawpi/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.GET("/users/:id", controller.GetUserDetailController)

	//Generating Token
	e.POST("/login", controller.LoginUsersController)

	//Login
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Authenticated Router for users
	eJwt.GET("/users/:id", controller.GetUserDetailController)
	eJwt.DELETE("/users/:id", controller.DeleteControllerAuthenticated)
	eJwt.PUT("/users/:id", controller.UpdateControllerAuthenticated)

	// Authenticated Router for books
	eJwt.POST("/books", controller.CreateBookController)
	eJwt.DELETE("/books/:id", controller.DeleteBookController)
	eJwt.PUT("/books/:id", controller.UpdateBookController)

	//Not Authenticated for users
	e.GET("/users", controller.GetManyController)
	e.POST("/users", controller.CreateController)
	e.DELETE("/users/:id", controller.DeleteController)
	e.PUT("/users/:id", controller.UpdateController)

	//Not Authenticated for books
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)

}
