package gateway

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) repository.IEventRepository {
	return &eventRepository{
		db: db,
	}
}

func (r *eventRepository) FindByID(id int64) (*model.Event, error) {
	event := model.Event{}
	if err := r.db.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Error(err)
		return nil, errors.New(fmt.Sprintf("db access error. id: %d", id))
	}
	return &event, nil
}

func (r *eventRepository) FindByIDs(ids []int64) ([]model.Event, error) {
	var events []model.Event
	if err := r.db.Where("id IN (?)", ids).Find(&events).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Error(err)
		return nil, ErrDBAccessError
	}
	return events, nil
}

func (r *eventRepository) Store(event model.Event) (*int64, error) {
	if err := r.db.Create(&event).Error; err != nil {
		return nil, err
	}
	return &event.ID, nil
}
