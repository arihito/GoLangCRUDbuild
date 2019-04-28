package usecase

import (
	"github.com/sminoeee/sample-app/go/adapter/gateway"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/domain/repository"
	"github.com/sminoeee/sample-app/go/external/db"
)

type (
	IDashboardUseCase interface {
		FindEventsByUserID(id int64) ([]model.Event, error)
		FindStudyGroupsByUserID(id int64) ([]model.StudyGroup, error)
	}

	DashboardUseCase struct {
		EventReserveRepository repository.IEventReserveRepository
		EventRepository        repository.IEventRepository
	}
)

func NewDashboardUseCase() IDashboardUseCase {
	return &DashboardUseCase{
		EventReserveRepository: gateway.NewEventReserveRepository(db.Conn),
		EventRepository:        gateway.NewEventRepository(db.Conn),
	}
}

// ユーザーが参加するイベントを取得する
func (uc *DashboardUseCase) FindEventsByUserID(id int64) ([]model.Event, error) {
	// EventReserves から自分が参加するイベントを引く
	reserves, err := uc.EventReserveRepository.FindEventReservesByUserID(id)
	if err != nil {
		return nil, err
	}

	// event IDs
	var eventIDs []int64
	for _, e := range reserves {
		eventIDs = append(eventIDs, e.ID)
	}

	// Event を取得する
	events, err := uc.EventRepository.FindByIDs(eventIDs)
	if err != nil {
		return nil, err
	}

	return events, nil
}

// ユーザーが参加しているグループを取得する
func (uc *DashboardUseCase) FindStudyGroupsByUserID(id int64) ([]model.StudyGroup, error) {
	return nil, nil
}
