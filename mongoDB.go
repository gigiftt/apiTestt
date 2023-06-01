package main

import (
	"github.com/apiTestt/config"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // e.GET("/", func(c echo.Context) error {
    //       return c.JSON(200, &echo.Map{"data": "Hello from Echo & mongoDB"})
    // })

	config.ConnectDB()

    e.Logger.Fatal(e.Start(":8080"))
}