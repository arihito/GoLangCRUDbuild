package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type studyGroupMemberRepository struct {
	db *gorm.DB
}

func NewStudyGroupMemberRepository(db *gorm.DB) repository.StudyGroupMemberRepository {
	return &studyGroupMemberRepository{
		db: db,
	}
}

// TODO 実装
