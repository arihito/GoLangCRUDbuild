package usecase

import (
	"github.com/jinzhu/gorm"
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
		repository.IUserRepository
	}
)

func NewUserUseCase(db *gorm.DB) IUserUseCase {
	return &UserUseCase{
		IUserRepository: gateway.NewUserRepository(db),
	}
}

func (uc *UserUseCase) FindByID(id int64) (*model.User, error) {
	if seminar, err := uc.IUserRepository.FindByID(id); err != nil {
		return nil, err
	} else {
		return seminar, nil
	}
}

func (uc *UserUseCase) Store(user model.User) (*int64, error) {
	// TODO バリデーション ...
	if id, err := uc.IUserRepository.Store(user); err != nil {
		return nil, err
	} else {
		return id, nil
	}
}
