package handler

import (
	goModel "github.com/jeevatkm/go-model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/adapter/handler/response"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/external/db"
	"github.com/sminoeee/sample-app/go/usecase"
	"github.com/sminoeee/sample-app/go/util"
	"net/http"
	"strconv"
	"time"
)

type (
	IUserHandler interface {
		Get(ctx echo.Context) error
		GetStudyGroups(ctx echo.Context) error
		Store(ctx echo.Context) error
	}

	UserHandler struct {
		UserUseCase       usecase.IUserUseCase
		StudyGroupUseCase usecase.IStudyGroupUseCase
	}

	userDetailResponse struct {
		ID          int64
		Name        string
		DisplayName *string
		CanceledAt  *time.Time
	}

	userStoreRequest struct {
		Name        string
		DisplayName *string
	}

	userStoreResponse struct {
		ID int64
	}
)

func NewUserHandler() IUserHandler {
	return &UserHandler{
		UserUseCase:       usecase.NewUserUseCase(db.Conn),
		StudyGroupUseCase: usecase.NewStudyGroupUseCase(),
	}
}

// GET /users/:id
func (h *UserHandler) Get(ctx echo.Context) error {
	userID := util.GetUserIDFromRequest(ctx)

	user, err := h.UserUseCase.FindByID(userID)
	if err != nil {
		log.Error(err)
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}

	if user == nil {
		return NewApplicationError(http.StatusNotFound, "not found.")
	}

	res := userDetailResponse{}
	if errs := goModel.Copy(&res, user); len(errs) > 0 {
		log.Error(errs)
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}

	return ctx.JSON(http.StatusOK, res)
}

// GET /users/:id/study_groups
func (h *UserHandler) GetStudyGroups(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	groups, err := h.StudyGroupUseCase.FindByUserID(int64(userID))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewStudyGroupResponses(groups))
}

// POST /users
func (h *UserHandler) Store(ctx echo.Context) error {
	req := userStoreRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Error(err)
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	user := model.User{}
	if errs := goModel.Copy(&user, req); len(errs) > 0 {
		log.Error(errs)
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	id, err := h.UserUseCase.Store(user);
	if err != nil {
		log.Error(err)
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}

	return ctx.JSON(http.StatusCreated, userStoreResponse{ID: *id})
}
