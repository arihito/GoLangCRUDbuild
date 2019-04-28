package model

import (
	"time"
)

type RaffleMethod int

const (
	RaffleMethodFirstCome RaffleMethod = iota // 先着
	RaffleMethodRaffle                        // 抽選
)

// イベント参加枠
type EventSeat struct {
	ID           int64
	EventID      int64
	Title        string
	MaximumLimit int
	ReservedNum  int
	Fee          int
	RaffleMethod RaffleMethod
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Event *Event
}

func (m *EventSeat) TableName() string {
	return "event_seats"
}
