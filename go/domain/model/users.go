package model

import (
	"time"
)

type User struct {
	ID          int64
	Name        string
	DisplayName *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *User) TableName() string {
	return "users"
}
