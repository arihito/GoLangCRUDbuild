package handler

import (
	goModel "github.com/jeevatkm/go-model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/domain/model"
	"github.com/sminoeee/sample-app/go/usecase/interfaces"
	"net/http"
	"strconv"
	"time"
)

type (
	userHandler struct {
		useCase interfaces.UserUseCase
	}

	userDetailResponse struct {
		ID          uint64
		Name        string
		DisplayName *string
		CanceledAt  *time.Time
	}

	userStoreRequest struct {
		Name        string
		DisplayName *string
	}

	userStoreResponse struct {
		ID uint64
	}
)

func NewUserHandler(useCase interfaces.UserUseCase) *userHandler {
	return &userHandler{
		useCase: useCase,
	}
}

// GET /users/:id
func (handler *userHandler) FindById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	user, err := handler.useCase.FindByID(uint64(id))
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

// POST /users
func (handler *userHandler) Store(ctx echo.Context) error {
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

	id, err := handler.useCase.Store(user);
	if err != nil {
		log.Error(err)
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}

	return ctx.JSON(http.StatusCreated, userStoreResponse{ID: *id})
}
