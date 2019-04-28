package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type eventSeatRepository struct {
	db *gorm.DB
}

func NewEventSeatRepository(db *gorm.DB) repository.IEventSeatRepository {
	return &eventSeatRepository{
		db: db,
	}
}

// TODO 実装
