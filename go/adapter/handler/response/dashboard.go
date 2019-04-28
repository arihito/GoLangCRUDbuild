package response

import (
	"github.com/sminoeee/sample-app/go/domain/model"
	"time"
)

type (
	DashboardResponse struct {
		Events      []Event
		StudyGroups []StudyGroup
	}

	// レスポンス用のイベント
	Event struct {
		ID           int64
		Title        string
		SubTitle     *string
		ImagePath    *string
		StudyGroupID *int64
		EventStart   time.Time
		EventEnd     time.Time
		ApplyStart   time.Time
		ApplyEnd     time.Time
		Summary      string
		UserID       int64
		Published    bool
	}

	// レスポンス用の勉強会グループ
	StudyGroup struct {
		ID        int64
		Name      string
		PageUrl   string
		UserID    int64
		Published bool
	}
)

func NewDashboardResponse(events []model.Event, groups []model.StudyGroup) *DashboardResponse {
	newEvent := func(e model.Event) Event {
		return Event{
			ID:           e.ID,
			Title:        e.Title,
			SubTitle:     e.SubTitle,
			ImagePath:    e.ImagePath,
			StudyGroupID: e.StudyGroupID,
			EventStart:   e.EventStart,
			EventEnd:     e.EventEnd,
			ApplyStart:   e.ApplyStart,
			ApplyEnd:     e.ApplyEnd,
			Summary:      e.Summary,
			UserID:       e.UserID,
			Published:    e.Published,
		}
	}
	var retEvents []Event
	for _, e := range events {
		retEvents = append(retEvents, newEvent(e))
	}

	newStudyGroup := func(g model.StudyGroup) StudyGroup {
		return StudyGroup{
			ID:        g.ID,
			Name:      g.Name,
			PageUrl:   g.PageUrl,
			UserID:    g.UserID,
			Published: g.Published,
		}
	}
	var retGroups []StudyGroup
	for _, g := range groups {
		retGroups = append(retGroups, newStudyGroup(g))
	}

	return &DashboardResponse{
		Events:      retEvents,
		StudyGroups: retGroups,
	}
}
