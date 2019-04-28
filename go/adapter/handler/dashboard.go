package handler

import (
	"github.com/labstack/echo"
	"github.com/sminoeee/sample-app/go/adapter/handler/response"
	"github.com/sminoeee/sample-app/go/usecase"
	"github.com/sminoeee/sample-app/go/util"
	"net/http"
)

type (
	IDashboardHandler interface {
		Get(c echo.Context) error
	}

	DashboardHandler struct {
		uc usecase.IDashboardUseCase
	}
)

func NewDashboardHandler() IDashboardHandler {
	return &DashboardHandler{
		uc: usecase.NewDashboardUseCase(),
	}
}

func (h *DashboardHandler) Get(c echo.Context) error {
	// request からユーザーIDを取得する
	userID := util.GetUserIDFromRequest(c)

	// 自分が参加するイベント
	events, err := h.uc.FindEventsByUserID(userID)
	if err != nil {
		return err
	}

	// 自分が所属するグループ
	groups, err := h.uc.FindStudyGroupsByUserID(userID)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		response.NewDashboardResponse(events, groups),
	)
}
