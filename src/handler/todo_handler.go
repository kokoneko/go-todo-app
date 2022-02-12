package handler

import (
	"net/http"
	"go-todo-app/usecase"
	"github.com/labstack/echo/v4"
)

func GetTodoList(c echo.Context) error {
	res, err := usecase.GetTodoList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, []string{})
	}

	return c.JSONPretty(http.StatusOK, res, " ")
}