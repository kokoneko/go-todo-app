package tests

import (
	"testing"
	"time"

	"github.com/kokoneko/go-todo-app/domain"
	"github.com/kokoneko/go-todo-app/infrastructure"
	"github.com/kokoneko/go-todo-app/usecase"
)

func TestDeleteTodoItem(t *testing.T) {
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

	initData := []*domain.TodoItem{
		todoItem,
	}

	for _, v := range initData {
		db.Create(v)
	}

	tests := []struct {
		name     string
		expected error
		isErr    bool
	}{
		{
			name:     "200: delete todo item",
			expected: nil,
			isErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := usecase.DeleteTodoItem(todoItem)

			if err != nil && tt.isErr {
				t.Errorf("err: %v, expected: %v", err, tt.expected)
			}
		})
	}
}
