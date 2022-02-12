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

func GetTodoById(id int) (*domain.TodoItem, error) {
	todo, err := repository.GetTodoById(id)

	if err != nil {
		log.Debug("get todo item error", err)
		return nil, err
	}

	return todo, nil
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

type UpdateTodoItemRequest struct {
	Title   string     `json:"title" validate:"required"`
	Memo    string     `json:"memo"`
	Expired *time.Time `json:"expired,omitempty"`
}

func UpdateTodoItem(todo *domain.TodoItem, p *UpdateTodoItemRequest) (*domain.TodoItem, error) {
	ut := map[string]interface{}{
		"title":   p.Title,
		"memo":    p.Memo,
		"expired": p.Expired,
	}

	if _, err := repository.UpdateTodoItem(todo, ut); err != nil {
		log.Debug("update todo item error", err)
		return nil, err
	}

	return todo, nil
}

func DeleteTodoItem(todo *domain.TodoItem) error {
	err := repository.DeleteTodoItem(todo)
	if err != nil {
		log.Debug("delete todo item error", err)
		return err
	}

	return nil
}
