package usecase

import (
	"time"

	"github.com/kokoneko/go-todo-app/domain"
	"github.com/kokoneko/go-todo-app/repository"

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
	ID      uint32     `json:"id"`
	Title   string     `json:"title"`
	Memo    string     `json:"memo"`
	Expired *time.Time `json:"expired"`
}

func GetTodoById(id int) (*GetTodoItemResponse, error) {
	todo, err := repository.GetTodoById(id)

	if err != nil {
		log.Debug("get todo item error", err)
		return nil, err
	}

	return &GetTodoItemResponse{
		ID:      todo.ID,
		Title:   todo.Title,
		Memo:    todo.Memo,
		Expired: todo.Expired,
	}, nil
}

type CreateTodoItemRequest struct {
	Title   string     `json:"title" validate:"required"`
	Memo    string     `json:"memo"`
	Expired *time.Time `json:"expired,omitempty"`
}

func CreateTodoItem(p *CreateTodoItemRequest) error {
	now := time.Now()
	todo := &domain.TodoItem{
		Title:     p.Title,
		Memo:      p.Memo,
		Expired:   p.Expired,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := repository.CreateTodoItem(todo)
	if err != nil {
		log.Debug("create todo item error", err)
		return err
	}

	return nil
}
