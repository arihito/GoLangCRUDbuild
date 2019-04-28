package response

import "github.com/sminoeee/sample-app/go/domain/model"

type (
	StudyGroupResponses struct {
		StudyGroups []StudyGroupResponse
	}

	StudyGroupResponse struct {
		ID        int64
		Title     string
		SubTitle  string
		PageUrl   string
		UserID    int64
		Published bool
	}

	CreateStudyGroupResponse struct {
		ID int64
	}
)

func NewStudyGroupResponse(m model.StudyGroup) StudyGroupResponse {
	return StudyGroupResponse{
		ID:        m.ID,
		Title:     m.Title,
		SubTitle:  m.SubTitle,
		PageUrl:   m.PageUrl,
		UserID:    m.UserID,
		Published: m.Published,
	}
}

func NewStudyGroupResponses(models []model.StudyGroup) StudyGroupResponses {
	var ret []StudyGroupResponse
	for _, sg := range models {
		ret = append(ret, NewStudyGroupResponse(sg))
	}
	return StudyGroupResponses{
		StudyGroups: ret,
	}
}
