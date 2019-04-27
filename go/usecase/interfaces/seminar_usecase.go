package interfaces

import "github.com/sminoeee/sample-app/go/domain/model"

type SeminarUseCase interface {
	FindSeminars() ([]model.Seminar, error)
	FindSeminarDetail(id uint64) (*model.Seminar, error)

	Store(seminar model.Seminar) (*uint64, error)

	// TODO 拡張
}
