package gateway

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/domain/repository"
)

type studyGroupRepository struct {
	db *gorm.DB
}

func NewStudyGroupRepository(db *gorm.DB) repository.IStudyGroupRepository {
	return &studyGroupRepository{
		db: db,
	}
}

func (r *studyGroupRepository) FindByID(id int64) (*model.StudyGroup, error) {
	var ret model.StudyGroup
	// select * from study_groups where id = :id
	if err := r.db.Where("id = ?", id).Find(&ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Error(err)
		return nil, ErrDBAccessError
	}
	return &ret, nil
}

func (r *studyGroupRepository) FindByIDs(ids []int64) ([]model.StudyGroup, error) {
	var ret []model.StudyGroup
	// select * from study_groups where id in (ids)
	if err := r.db.Where("id IN (?)", ids).Find(&ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Error(err)
		return nil, ErrDBAccessError
	}
	return nil, nil
}

// グループを登録
func (r *studyGroupRepository) Create(studyGroup model.StudyGroup) (*int64, error) {
	if err := r.db.Create(&studyGroup).Error; err != nil {
		log.Error(err)
		return nil, ErrDBAccessError
	}
	return &studyGroup.ID, nil
}

// グループを更新
func (r *studyGroupRepository) Update(id int64, sg model.StudyGroup) error {
	// golang のゼロ値は無視されるためカラム指定で更新する
	updateMap := map[string]interface{}{
		"title":      sg.Title,
		"sub_title":  sg.SubTitle,
		"image_path": sg.ImagePath,
		"page_url":   sg.PageUrl,
		"published":  sg.Published,
	}

	// update study_groups set title = ?, ... where id = :id
	if err := r.db.Model(model.StudyGroup{}).Where("id = ?", id).Updates(updateMap).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		log.Error(err)
		return ErrDBAccessError
	}
	return nil
}

// グループを削除
func (r *studyGroupRepository) Delete(id int64) error {
	// delete from study_groups where id = :id
	if err := r.db.Where("id = ?", id).Delete(model.StudyGroup{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		log.Error(err)
		return ErrDBAccessError
	}
	return nil
}
