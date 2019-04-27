package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/adapter/gateway"
	"github.com/sminoeee/sample-app/go/domain/model"
	interfaces2 "github.com/sminoeee/sample-app/go/domain/repository/interfaces"
	"github.com/sminoeee/sample-app/go/usecase/interfaces"
)

type userUseCase struct {
	interfaces2.UserRepository
}

func NewUserUseCase(db *gorm.DB) interfaces.UserUseCase {
	return &userUseCase{
		UserRepository: gateway.NewUserRepository(db),
	}
}

func (uc *userUseCase) FindByID(id uint64) (*model.User, error) {
	if seminar, err := uc.UserRepository.FindByID(id); err != nil {
		log.Error(err)
		return nil, ErrFailedDbAccess
	} else {
		return seminar, nil
	}
}

func (uc *userUseCase) Store(user model.User) (*uint64, error) {
	// TODO バリデーション ...
	if id, err := uc.UserRepository.Store(user); err != nil {
		log.Error(err)
		return nil, ErrFailedDbAccess
	} else {
		return id, nil
	}
}
