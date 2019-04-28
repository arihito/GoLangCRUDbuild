package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type studyGroupRepository struct {
	db *gorm.DB
}

func NewStudyGroupRepository(db *gorm.DB) repository.StudyGroupRepository {
	return &studyGroupRepository{
		db: db,
	}
}

// TODO 実装
