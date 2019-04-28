package model

import "time"

type EventTag struct {
	ID        int64
	EventID   int64
	TagID     int64
	CreatedAt time.Time
	UpdatedAt time.Time

	Event *Event
	Tag   *Tag
}

func (m *EventTag) TableName() string {
	return "event_tags"
}
