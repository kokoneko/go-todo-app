package usecase

import (
	"time"

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

type GetTodoItemResponse struct {
	ID 		uint32 	  `json:"id"`
	Title 	string 	  `json:"title"`
	Memo 	string 	  `json:"memo"`
	Expired time.Time `json:"expired"`
}

func GetTodoById(id int) (*GetTodoItemResponse, error) {
	todo, err := repository.GetTodoById(id)

	if err != nil {
		log.Debug("get todo item error", err)
		return nil, err
	}

	return &GetTodoItemResponse{
		ID:	todo.ID,
		Title: todo.Title,
		Memo: todo.Memo,
		Expired: todo.Expired,
	}, nil
}