package usecase

import (
	"go-todo-app/repository"
	"go-todo-app/domain"

	"github.com/labstack/gommon/log"
)

type GetTodoListResponse struct {
	List []*domain.TodoItem `json:"lists"`
}

func GetTodoList() (*GetTodoListResponse, error) {
	todoList, err := repository.GetTodoList()

	if err != nil {
		log.Debug("get todo list error", err)
		return nil, err
	}

	return &GetTodoListResponse{List: todoList}, nil
}