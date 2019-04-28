package model

import "time"

type EventReserve struct {
	ID          int64
	UserID      int64
	EventID     int64
	EventSeatID int64
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User      *User
	Event     *Event
	EventSeat *EventSeat
}

func (m *EventReserve) TableName() string {
	return "event_reserves"
}
