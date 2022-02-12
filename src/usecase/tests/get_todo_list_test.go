package tests

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kokoneko/go-todo-app/domain"
	"github.com/kokoneko/go-todo-app/infrastructure"
	"github.com/kokoneko/go-todo-app/usecase"
)

func TestGetTodoList(t *testing.T) {
	db := infrastructure.ConnectDB()

	defer truncateTable(db)

	jst, _ := time.LoadLocation("Asia/Tokyo")
	tm := time.Now().In(jst)
	now := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, jst)

	todoItem1 := &domain.TodoItem{
		ID:        1,
		Title:     "test title1",
		Memo:      "test memo1",
		Expired:   &now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	todoItem2 := &domain.TodoItem{
		ID:        2,
		Title:     "test title2",
		Memo:      "test memo2",
		Expired:   &now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	initData := []*domain.TodoItem{
		todoItem1,
		todoItem2,
	}

	for _, v := range initData {
		db.Create(v)
	}

	tests := []struct {
		name     string
		expected usecase.GetTodoListResponse
		isErr    bool
	}{
		{
			name: "200: get todo list",
			expected: usecase.GetTodoListResponse{
				List: initData,
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := usecase.GetTodoList()

			if err != nil && tt.isErr {
				t.Errorf("err: %v, expected: %v", err, tt.expected)
			}

			for i, got := range ret.List {
				if d := cmp.Diff(got, initData[i]); len(d) != 0 {
					t.Errorf("differs: (-got +want)\n%s", d)
				}
			}
		})
	}
}
