package usecase

import (
	"errors"
	"fmt"
	"github.com/sminoeee/sample-app/go/adapter/gateway"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/domain/repository"
	db2 "github.com/sminoeee/sample-app/go/external/db"
)

type (
	IStudyGroupUseCase interface {
		FindByID(id int64) (*model.StudyGroup, error)
		FindByUserID(userID int64) ([]model.StudyGroup, error)

		Create(userID int64, title string) (*int64, error)
		Update(userID int64, groupID int64, studyGroup model.StudyGroup) error

		Delete(userID int64, groupID int64) error
	}

	StudyGroupUseCase struct {
		StudyGroupMemberRepo repository.IStudyGroupMemberRepository
		StudyGroupRepo       repository.IStudyGroupRepository
	}
)

func NewStudyGroupUseCase() IStudyGroupUseCase {
	return &StudyGroupUseCase{
		StudyGroupMemberRepo: gateway.NewStudyGroupMemberRepository(db2.Conn),
		StudyGroupRepo:       gateway.NewStudyGroupRepository(db2.Conn),
	}
}

// 勉強会グループを取得する
func (uc *StudyGroupUseCase) FindByID(id int64) (*model.StudyGroup, error) {
	group, err := uc.StudyGroupRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// ユーザーが所属する勉強会グループを取得する
func (uc *StudyGroupUseCase) FindByUserID(userID int64) ([]model.StudyGroup, error) {
	groupMembers, err := uc.StudyGroupMemberRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	// group ids
	var groupIDs []int64
	for _, g := range groupMembers {
		groupIDs = append(groupIDs, g.ID)
	}

	groups, err := uc.StudyGroupRepo.FindByIDs(groupIDs)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// 勉強会グループを作成する
func (uc *StudyGroupUseCase) Create(userID int64, title string) (*int64, error) {
	sg := model.StudyGroup{
		Title:     title,
		UserID:    userID, // 主催者は作成者
		Published: false,  // デフォルトは非公開
	}

	groupID, err := uc.StudyGroupRepo.Create(sg)
	if err != nil {
		return nil, err
	}

	return groupID, nil
}

// 勉強会グループを更新する
// FIXME: Application Error と HttpError は分けたい
func (uc *StudyGroupUseCase) Update(userID int64, groupID int64, studyGroup model.StudyGroup) error {
	group, err := uc.FindByID(groupID)
	if err != nil {
		return err
	}

	if group == nil {
		// 更新対象の存在チェック
		return errors.New(fmt.Sprintf("StudyGroup does not found. id: %d", groupID))
	}

	if group.UserID != userID {
		// 他人のグループは更新できない
		return errors.New(fmt.Sprintf("Permission error."))
	}

	if err := uc.StudyGroupRepo.Update(groupID, studyGroup); err != nil {
		return err
	}

	return nil
}

// グループを削除する
func (uc *StudyGroupUseCase) Delete(userID int64, groupID int64) error {
	group, err := uc.FindByID(groupID)
	if err != nil {
		return err
	}

	if group == nil {
		// 更新対象の存在チェック
		return errors.New(fmt.Sprintf("StudyGroup does not found. id: %d", groupID))
	}

	if group.UserID != userID {
		// 他人のグループは削除できない
		return errors.New(fmt.Sprintf("Permission error."))
	}

	if err := uc.StudyGroupRepo.Delete(groupID); err != nil {
		return err
	}

	return nil
}
