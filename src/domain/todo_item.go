package domain

import (
	"time"
)

type TodoItem struct {
	ID        uint32     `json:"id" gorm:"primary_key"`
	Title     string     `json:"title"`
	Memo      string     `json:"memo"`
	Expired   *time.Time `json:"expired"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
