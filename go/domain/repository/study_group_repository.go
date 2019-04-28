package repository

import "github.com/sminoeee/sample-app/go/domain/model"

type IStudyGroupRepository interface {
	FindByID(id int64) (*model.StudyGroup, error)
	FindByIDs(ids []int64) ([]model.StudyGroup, error)

	// Store でいいかも
	Create(studyGroup model.StudyGroup) (*int64, error)

	Update(id int64, studyGroup model.StudyGroup) error

	Delete(id int64) error
}
