package gateway

import (
	"github.com/jinzhu/gorm"
)

type seminarSeatRepository struct {
	db *gorm.DB
}

func NewSeminarSeatRepository(db *gorm.DB) *seminarSeatRepository {
	return &seminarSeatRepository{
		db: db,
	}
}

// TODO 実装
