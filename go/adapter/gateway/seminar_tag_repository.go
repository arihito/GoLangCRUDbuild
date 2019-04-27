package gateway

import (
	"github.com/jinzhu/gorm"
)

type seminarTagRepository struct {
	db *gorm.DB
}

func NewSeminarTagRepository(db *gorm.DB) *seminarTagRepository {
	return &seminarTagRepository{
		db: db,
	}
}

// TODO 実装
