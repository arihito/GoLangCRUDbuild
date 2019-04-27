package model

import (
	"time"
)

type SeminarSeat struct {
	ID             uint64
	SeminarId      int64
	Summary        string
	RaffleDate     time.Time
	MaximumLimit   int
	ReservationNum int
}
