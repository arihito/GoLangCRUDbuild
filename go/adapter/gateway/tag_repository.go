package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) repository.TagRepository {
	return &tagRepository{
		db: db,
	}
}
