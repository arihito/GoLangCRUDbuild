package repository

import "github.com/sminoeee/sample-app/go/domain/model"

type UserRepository interface {
	FindByID(id int64) (*model.User, error)
	Store(user model.User) (*int64, error)
}
