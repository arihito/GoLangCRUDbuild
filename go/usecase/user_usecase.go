package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/adapter/gateway"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type (
	IUserUseCase interface {
		FindByID(id int64) (*model.User, error)
		Store(user model.User) (*int64, error)
	}

	UserUseCase struct {
		repository.UserRepository
	}
)

func NewUserUseCase(db *gorm.DB) IUserUseCase {
	return &UserUseCase{
		UserRepository: gateway.NewUserRepository(db),
	}
}

func (uc *UserUseCase) FindByID(id int64) (*model.User, error) {
	if seminar, err := uc.UserRepository.FindByID(id); err != nil {
		log.Error(err)
		return nil, ErrFailedDbAccess
	} else {
		return seminar, nil
	}
}

func (uc *UserUseCase) Store(user model.User) (*int64, error) {
	// TODO バリデーション ...
	if id, err := uc.UserRepository.Store(user); err != nil {
		log.Error(err)
		return nil, ErrFailedDbAccess
	} else {
		return id, nil
	}
}
