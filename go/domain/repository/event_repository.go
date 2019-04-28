package repository

import "github.com/sminoeee/sample-app/go/domain/model"

type IEventRepository interface {
	FindByID(id int64) (*model.Event, error)
	FindByIDs(ids []int64) ([]model.Event, error)

	Store(seminar model.Event) (*int64, error)
}
