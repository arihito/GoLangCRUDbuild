package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type EventReserveRepository struct {
	db *gorm.DB
}

func NewEventReserveRepository(db *gorm.DB) repository.IEventReserveRepository {
	return &EventReserveRepository{
		db: db,
	}
}

func (r *EventReserveRepository) FindEventReservesByUserID(id int64) ([]model.EventReserve, error) {
	var ret []model.EventReserve

	// select * from event_reserves where user_id = :id
	if err := r.db.Where("user_id", id).Find(&ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Error(err)
		return nil, ErrDBAccessError
	}
	return ret, nil
}
