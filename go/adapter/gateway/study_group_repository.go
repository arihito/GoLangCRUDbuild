package gateway

import (
	"github.com/jinzhu/gorm"
)

type studyGroupRepository struct {
	db *gorm.DB
}

func NewStudyGroupRepository(db *gorm.DB) *studyGroupRepository {
	return &studyGroupRepository{
		db: db,
	}
}

// TODO 実装
