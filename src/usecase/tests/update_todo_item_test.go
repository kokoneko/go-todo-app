package tests

import (
	"testing"
	"time"

	"github.com/kokoneko/go-todo-app/domain"
	"github.com/kokoneko/go-todo-app/infrastructure"
	"github.com/kokoneko/go-todo-app/repository"
	"github.com/kokoneko/go-todo-app/usecase"
)

func TestUpdateTodoItem(t *testing.T) {
	db := infrastructure.ConnectDB()

	defer truncateTable(db)

	jst, _ := time.LoadLocation("Asia/Tokyo")
	tm := time.Now().In(jst)
	now := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, jst)

	todoItem := &domain.TodoItem{
		ID:        1,
		Title:     "test title",
		Memo:      "test memo",
		Expired:   &now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	updateTodo := &usecase.UpdateTodoItemRequest{
		Title:   "change.title",
		Memo:    "change memo",
		Expired: &now,
	}

	initData := []*domain.TodoItem{
		todoItem,
	}

	for _, v := range initData {
		db.Create(v)
	}

	tests := []struct {
		name     string
		expected *domain.TodoItem
		request  *usecase.UpdateTodoItemRequest
		isErr    bool
	}{
		{
			name: "200: update todo item",
			expected: &domain.TodoItem{
				Title:     updateTodo.Title,
				Memo:      updateTodo.Memo,
				Expired:   updateTodo.Expired,
				CreatedAt: now,
				UpdatedAt: now,
			},
			request: updateTodo,
			isErr:   false,
		},
	}

	var layout = "2006-01-02 15:04:05"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := usecase.UpdateTodoItem(todoItem, updateTodo)

			if err != nil && tt.isErr {
				t.Errorf("err: %v, expected: %v", err, tt.expected)
			}

			res, _ := repository.GetTodoById(1)
			if res.Title != tt.request.Title {
				t.Errorf("err: %s, expected Title: %s", res.Title, tt.request.Title)
			}

			if res.Memo != tt.request.Memo {
				t.Errorf("err: %s, expected Memo: %s", res.Memo, tt.request.Memo)
			}

			if res.Expired.Format(layout) != tt.request.Expired.Format(layout) {
				t.Errorf("err: %s, expected Expired: %s", res.Expired, tt.request.Expired)
			}
		})
	}
}
