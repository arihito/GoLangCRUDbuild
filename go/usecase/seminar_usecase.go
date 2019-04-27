package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/adapter/gateway"
	"github.com/sminoeee/sample-app/go/domain/model"
	interfaces2 "github.com/sminoeee/sample-app/go/domain/repository/interfaces"
	"github.com/sminoeee/sample-app/go/usecase/interfaces"
)

type seminarUseCase struct {
	interfaces2.SeminarRepository
}

func NewSeminarUseCase(db *gorm.DB) interfaces.SeminarUseCase {
	return &seminarUseCase{
		SeminarRepository: gateway.NewSeminarRepository(db),
	}
}

func (uc *seminarUseCase) FindSeminars() ([]model.Seminar, error) {
	// TODO 実装
	return nil, nil
}

func (uc *seminarUseCase) FindSeminarDetail(id uint64) (*model.Seminar, error) {
	if seminar, err := uc.SeminarRepository.Find(id); err != nil {
		log.Error(err)
		return nil, ErrFailedDbAccess
	} else {
		return seminar, nil
	}
}

func (uc *seminarUseCase) Store(seminar model.Seminar) (*uint64, error) {
	// TODO バリデーション ...
	if id, err := uc.SeminarRepository.Store(seminar); err != nil {
		log.Error(err)
		return nil, ErrFailedDbAccess
	} else {
		return id, nil
	}
}
