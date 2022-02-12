package tests

import (
	"testing"
	"time"

	"github.com/kokoneko/go-todo-app/infrastructure"
	"github.com/kokoneko/go-todo-app/repository"
	"github.com/kokoneko/go-todo-app/usecase"
)

func TestCreateTodoItem(t *testing.T) {
	db := infrastructure.ConnectDB()

	defer truncateTable(db)

	jst, _ := time.LoadLocation("Asia/Tokyo")
	tm := time.Now().In(jst)
	now := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, jst)

	tests := []struct {
		name     string
		expected error
		request  *usecase.CreateTodoItemRequest
		isErr    bool
	}{
		{
			name:     "200: create todo item",
			expected: nil,
			request: &usecase.CreateTodoItemRequest{
				Title:   "test title",
				Memo:    "test memo",
				Expired: &now,
			},
			isErr: false,
		},
	}

	var cnt int
	var layout = "2006-01-02 15:04:05"
	for _, tt := range tests {
		cnt++
		t.Run(tt.name, func(t *testing.T) {
			err := usecase.CreateTodoItem(tt.request)

			if err != nil && tt.isErr {
				t.Errorf("err: %v, expected: %v", err, tt.expected)
			}

			res, _ := repository.GetTodoById(cnt)
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
