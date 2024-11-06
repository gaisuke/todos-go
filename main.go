package main

import (
	"fmt"
	"todos-go/pkg/db/postgres"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	postgres.InitDB()

	fmt.Println("server running localhost:8080")
	e.Logger.Fatal(e.Start("localhost:8080"))
}
