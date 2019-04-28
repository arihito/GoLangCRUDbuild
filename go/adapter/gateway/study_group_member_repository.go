package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type studyGroupMemberRepository struct {
	db *gorm.DB
}

func NewStudyGroupMemberRepository(db *gorm.DB) repository.IStudyGroupMemberRepository {
	return &studyGroupMemberRepository{
		db: db,
	}
}

func (r studyGroupMemberRepository) FindByUserID(userID int64) ([]model.StudyGroupMember, error) {
	var ret []model.StudyGroupMember
	if err := r.db.Where("user_id = ?", userID).Find(&ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Error(err)
		return nil, ErrDBAccessError
	}
	return ret, nil
}

// TODO 実装
