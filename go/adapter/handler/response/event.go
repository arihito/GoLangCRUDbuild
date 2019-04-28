package response

import (
	"github.com/sminoeee/sample-app/go/domain/model"
	"time"
)

type (
	EventResponses struct {
		Events []EventResponse
	}

	// レスポンス用のイベント
	EventResponse struct {
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
)

func NewEventResponse(e model.Event) EventResponse {
	return EventResponse{
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

func NewEventResponses(models []model.Event) EventResponses {
	var ret []EventResponse
	for _, e := range models {
		ret = append(ret, NewEventResponse(e))
	}
	return EventResponses{
		Events: ret,
	}
}
