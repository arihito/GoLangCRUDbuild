package repository

import (
	"github.com/sminoeee/sample-app/go/domain/model"
)

type IEventReserveRepository interface {
	FindEventReservesByUserID(id int64) ([]model.EventReserve, error)
}
