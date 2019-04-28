package model

import "time"

type Event struct {
	ID           int64
	Title        string
	SubTitle     *string
	ImagePath    *string
	StudyGroupID *int64
	EventStart   time.Time
	EventEnd     time.Time
	ApplyStart   time.Time
	ApplyEnd     time.Time
	Summary      string
	UserID       int64
	Published    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time

	StudyGroup *StudyGroup
	User       *User
}

func (m *Event) TableName() string {
	return "events"
}
