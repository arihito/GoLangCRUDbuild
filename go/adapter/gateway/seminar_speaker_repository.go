package gateway

import (
	"github.com/jinzhu/gorm"
)

type seminarSpeakerRepository struct {
	db *gorm.DB
}

func NewSeminarSpeakerRepository(db *gorm.DB) *seminarSpeakerRepository {
	return &seminarSpeakerRepository{
		db: db,
	}
}

// TODO 実装
