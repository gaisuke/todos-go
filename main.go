package main

import (
	"fmt"
	"todos-go/controllers"
	"todos-go/handlers"
	"todos-go/pkg/db/postgres"
	"todos-go/services"

	"github.com/labstack/echo/v4"
)

func main() {
	postgres.InitDB()

	db := postgres.DB

	listService := services.NewListService(db)
	sublistService := services.NewSublistService(db)
	listController := controllers.NewListController(listService)
	sublistController := controllers.NewSublistController(sublistService)

	handler := handlers.NewTodoHandler(listController, sublistController)

	e := echo.New()
	registerRoutes(e, handler)

	// for testing only
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(200, "Hello, World!")
	// })

	fmt.Println("server running localhost:8080")
	e.Logger.Fatal(e.Start("localhost:8080"))
}

func registerRoutes(e *echo.Echo, handler *handlers.TodoHandler) {
	e.GET("/lists", handler.GetLists)
	e.GET("/lists/:id", handler.GetListByID)
	e.GET("/lists/:id/sublists", handler.GetSublistsByListID)
	e.GET("/sublists/:id", handler.GetSublistByID)
	e.POST("/lists", handler.CreateList)
	e.POST("/lists/:id/sublists", handler.CreateSublist)
	e.PUT("/lists/:id", handler.UpdateList)
	e.PUT("/sublists/:id", handler.UpdateSublist)
	e.DELETE("/lists/:id", handler.DeleteList)
	e.DELETE("/sublists/:id", handler.DeleteSublist)
}
