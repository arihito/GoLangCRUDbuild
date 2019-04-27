package model

import "github.com/sminoeee/sample-app/go/domain/enum"

type StudyGroupMember struct {
	ID           uint64
	StudyGroupId int64
	UserId       int64
	Kind         enum.Kind
}
