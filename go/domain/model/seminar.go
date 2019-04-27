package model

import (
	"time"
)

type Seminar struct {
	ID        uint64
	Title     string
	SubTitle  *string
	EventDate time.Time
	Summary   string
	Organizer int64
	Start     time.Time
	End       time.Time
	Location  string
}
