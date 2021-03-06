package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/kokoneko/go-todo-app/usecase"
	"github.com/labstack/echo/v4"
)

func GetTodoList(c echo.Context) error {
	res, err := usecase.GetTodoList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, []string{})
	}

	return c.JSONPretty(http.StatusOK, res, " ")
}

func GetTodoById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := usecase.GetTodoById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSONPretty(http.StatusOK, res, " ")
}

func CreateTodoItem(c echo.Context) error {
	p := new(usecase.CreateTodoItemRequest)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	if err := usecase.CreateTodoItem(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, []string{})
}

func UpdateTodoItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := usecase.GetTodoById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	p := new(usecase.UpdateTodoItemRequest)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	res, err := usecase.UpdateTodoItem(todo, p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func DeleteTodoItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := usecase.GetTodoById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := usecase.DeleteTodoItem(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, []string{})
}
