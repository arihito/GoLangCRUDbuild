package gateway

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/domain/model"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByID(id uint64) (*model.User, error) {
	user := model.User{}
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Error(err)
		return nil, errors.New(fmt.Sprintf("db access error. id: %d", id))
	}
	return &user, nil
}

func (r *userRepository) Store(user model.User) (*uint64, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user.ID, nil
}
