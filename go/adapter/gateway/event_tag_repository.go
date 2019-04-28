package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type eventTagRepository struct {
	db *gorm.DB
}

func NewEventTagRepository(db *gorm.DB) repository.EventTagRepository {
	return &eventTagRepository{
		db: db,
	}
}

// TODO 実装
