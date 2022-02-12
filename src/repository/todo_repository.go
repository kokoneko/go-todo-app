package repository

import (
    "go-todo-app/domain"
    "go-todo-app/infrastructure"
)

func GetTodoList() ([]*domain.TodoItem, error) {
    db := infrastructure.ConnectDB()

    var todoList []*domain.TodoItem
    if err := db.Find(&todoList).Error; err != nil {
        return nil, err
    }

    return todoList, nil
}