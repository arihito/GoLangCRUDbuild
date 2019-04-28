package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/adapter/handler/request"
	"github.com/sminoeee/sample-app/go/adapter/handler/response"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/usecase"
	"github.com/sminoeee/sample-app/go/util"
	"net/http"
	"strconv"
)

type (
	IStudyGroupHandler interface {
		Get(ctx echo.Context) error
		Create(ctx echo.Context) error
		Update(ctx echo.Context) error
		Delete(ctx echo.Context) error
	}

	StudyGroupHandler struct {
		StudyGroupUseCase usecase.IStudyGroupUseCase
	}
)

func NewStudyGroupHandler() IStudyGroupHandler {
	return &StudyGroupHandler{
		StudyGroupUseCase: usecase.NewStudyGroupUseCase(),
	}
}

// グループを取得する
func (h *StudyGroupHandler) Get(ctx echo.Context) error {
	groupID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	group, err := h.StudyGroupUseCase.FindByID(int64(groupID))
	if err != nil {
		return err
	}

	if group == nil {
		return NewApplicationError(http.StatusNotFound, "not found.")
	}

	return ctx.JSON(
		http.StatusOK,
		response.NewStudyGroupResponse(*group),
	)
}

// グループ作成
func (h *StudyGroupHandler) Create(ctx echo.Context) error {
	req := request.CreateStudyGroupRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Error(err)
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	userID := util.GetUserIDFromRequest(ctx)

	groupID, err := h.StudyGroupUseCase.Create(userID, req.Title)
	if err != nil {
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}

	return ctx.JSON(
		http.StatusCreated,
		response.CreateStudyGroupResponse{ID: *groupID},
	)
}

// グループ更新
func (h *StudyGroupHandler) Update(ctx echo.Context) error {
	// パスID
	groupID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	// リクエストボディ
	req := request.UpdateStudyGroupRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Error(err)
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	userID := util.GetUserIDFromRequest(ctx)

	// 更新要求
	updateRequest := model.StudyGroup{
		Title:     req.Title,
		SubTitle:  req.SubTitle,
		ImagePath: req.ImagePath,
		PageUrl:   req.PageUrl,
		Published: req.Published,
	}

	if err := h.StudyGroupUseCase.Update(userID, int64(groupID), updateRequest); err != nil {
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}

	return ctx.NoContent(http.StatusNoContent)
}

// グループ削除
func (h *StudyGroupHandler) Delete(ctx echo.Context) error {
	groupID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	userID := util.GetUserIDFromRequest(ctx)

	if err := h.StudyGroupUseCase.Delete(userID, int64(groupID)); err != nil {
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}

	return ctx.NoContent(http.StatusNoContent)
}
