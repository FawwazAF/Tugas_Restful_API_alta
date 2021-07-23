package main

import (
	"example/pawpi/config"
	"example/pawpi/routes"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDb()
	config.InitPort()
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
	// e.Logger.Fatal(e.Start(":8080"))
}
