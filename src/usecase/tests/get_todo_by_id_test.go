package tests

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kokoneko/go-todo-app/domain"
	"github.com/kokoneko/go-todo-app/infrastructure"
	"github.com/kokoneko/go-todo-app/usecase"
	"gorm.io/gorm"
)

func truncateTable(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE todo_items")
}

func TestGetTodoById(t *testing.T) {
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
		expected *domain.TodoItem
		isErr    bool
	}{
		{
			name: "200: get todo by id",
			expected: &domain.TodoItem{
				ID:        todoItem.ID,
				Title:     todoItem.Title,
				Memo:      todoItem.Memo,
				Expired:   todoItem.Expired,
				CreatedAt: todoItem.CreatedAt,
				UpdatedAt: todoItem.UpdatedAt,
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := usecase.GetTodoById(1)

			if err != nil && tt.isErr {
				t.Errorf("err: %v, expected: %v", err, tt.expected)
			}

			if d := cmp.Diff(ret, tt.expected); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
