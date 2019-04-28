package response

import (
	"github.com/sminoeee/sample-app/go/domain/model"
)

type (
	DashboardResponse struct {
		Events      EventResponses
		StudyGroups StudyGroupResponses
	}
)

func NewDashboardResponse(events []model.Event, groups []model.StudyGroup) *DashboardResponse {
	return &DashboardResponse{
		Events:      NewEventResponses(events),
		StudyGroups: NewStudyGroupResponses(groups),
	}
}
