package model

import (
	"time"
)

type User struct {
	ID          uint64
	Name        string
	DisplayName *string
	CanceledAt  *time.Time
}
