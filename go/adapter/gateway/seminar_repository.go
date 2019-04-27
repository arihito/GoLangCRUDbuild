package gateway

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/domain/model"
)

type seminarRepository struct {
	db *gorm.DB
}

func NewSeminarRepository(db *gorm.DB) *seminarRepository {
	return &seminarRepository{
		db: db,
	}
}

func (r *seminarRepository) FindSeminars() ([]model.Seminar, error) {
	// TODO 実装
	return nil, nil
}

func (r *seminarRepository) Find(id uint64) (*model.Seminar, error) {
	seminar := model.Seminar{}
	// SELECT * FROM users WHERE id = {:id}
	if err := r.db.First(&seminar, id).Error; err != nil {
		// https://github.com/jinzhu/gorm/blob/master/errors.go
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Error(err)
		return nil, errors.New(fmt.Sprintf("db access error. id: %d", id))
	}
	return &seminar, nil
}

func (r *seminarRepository) Store(seminar model.Seminar) (*uint64, error) {
	if err := r.db.Create(&seminar).Error; err != nil {
		return nil, err
	}
	return &seminar.ID, nil
}
