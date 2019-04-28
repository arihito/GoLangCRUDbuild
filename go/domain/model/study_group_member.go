package model

import "time"

type MemberKind int // メンバー種別

const (
	MemberKindAdministrator MemberKind = iota // 管理者
	MemberKindMember                          // メンバー
)

type StudyGroupMember struct {
	ID           int64
	StudyGroupID int64
	UserID       int64
	MemberKind   MemberKind
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *StudyGroupMember) TableName() string {
	return "study_group_members"
}
