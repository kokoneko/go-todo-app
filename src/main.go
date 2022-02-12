package main

import (
	"net/http"

	"github.com/kokoneko/go-todo-app/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", hello)

	e.GET("/todo-list", handler.GetTodoList)
	e.GET("/todo-list/:id", handler.GetTodoById)
	e.POST("/todo-item", handler.CreateTodoItem)
	e.PUT("/todo-item/:id", handler.UpdateTodoItem)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
