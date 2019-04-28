package repository

import "github.com/sminoeee/sample-app/go/domain/model"

type IStudyGroupMemberRepository interface {
	FindByUserID(userID int64) ([]model.StudyGroupMember, error)
}
