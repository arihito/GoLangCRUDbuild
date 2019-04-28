package model

import "time"

type StudyGroup struct {
	ID        int64
	Name      string
	PageUrl   string
	UserID    int64
	Published bool
	CreatedAt time.Time
	UpdatedAt time.Time

	User *User
}

func (m *StudyGroup) TableName() string {
	return "study_groups"
}
