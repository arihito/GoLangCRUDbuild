package model

import "time"

type Tag struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Tag) TableName() string {
	return "tags"
}
