package gateway

import "github.com/jinzhu/gorm"

type reserveRepository struct {
	db *gorm.DB
}

func NewReserveRepository(db *gorm.DB) *reserveRepository {
	return &reserveRepository{
		db: db,
	}
}

// TODO 実装
