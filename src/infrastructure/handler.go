package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodoList(c echo.Context) error {
	return c.JSON(http.StatusOK, "todo")
}