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
	seminarHandler struct {
		UseCase interfaces.SeminarUseCase
	}

	DetailResponse struct {
		ID        uint64
		Title     string
		SubTitle  *string
		EventDate time.Time
		Summary   string
		Organizer int64
		Start     time.Time
		End       time.Time
		Location  string
	}

	StoreRequest struct {
		Title     string    `json:"title"`
		SubTitle  *string   `json:"subtitle"`
		EventDate time.Time `json:"eventdate"`
		Summary   string    `json:"summary"`
		Organizer int64     `json:"organizer"`
		Start     time.Time `json:"start"`
		End       time.Time `json:"end"`
		Location  string    `json:"location"`
	}

	StoreResponse struct {
		ID uint64
	}
)

func NewSeminarHandler(uc interfaces.SeminarUseCase) *seminarHandler {
	return &seminarHandler{
		UseCase: uc,
	}
}

// GET /seminars
func (handler *seminarHandler) FindSeminars(ctx echo.Context) error {
	// TODO: 実装
	return nil
}

// GET /seminars/:id
func (handler *seminarHandler) FindById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	seminar, err := handler.UseCase.FindSeminarDetail(uint64(id))
	if err != nil {
		log.Error(err)
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}

	if seminar == nil {
		return NewApplicationError(http.StatusNotFound, "not found.")
	}

	res := DetailResponse{}
	if errs := goModel.Copy(&res, seminar); len(errs) > 0 {
		log.Error(errs)
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}
	return ctx.JSON(http.StatusOK, res)
}

// POST /seminars
func (handler *seminarHandler) Store(ctx echo.Context) error {
	req := StoreRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Error(err)
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	seminar := model.Seminar{}
	if errs := goModel.Copy(&seminar, req); len(errs) > 0 {
		log.Error(errs)
		return NewApplicationError(http.StatusBadRequest, "bad request.")
	}

	id, err := handler.UseCase.Store(seminar);
	if err != nil {
		log.Error(err)
		return NewApplicationError(http.StatusInternalServerError, "server error.")
	}

	return ctx.JSON(http.StatusCreated, StoreResponse{ID: *id})
}
