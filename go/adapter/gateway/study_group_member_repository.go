package gateway

import (
	"github.com/jinzhu/gorm"
)

type studyGroupMemberRepository struct {
	db *gorm.DB
}

func NewStudyGroupMemberRepository(db *gorm.DB) *studyGroupMemberRepository {
	return &studyGroupMemberRepository{
		db: db,
	}
}

// TODO 実装
