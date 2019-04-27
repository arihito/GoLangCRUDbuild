package interfaces

import "github.com/sminoeee/sample-app/go/domain/model"

type SeminarRepository interface {
	FindSeminars() ([]model.Seminar, error)
	Find(id uint64) (*model.Seminar, error)

	Store(seminar model.Seminar) (*uint64, error)
}
