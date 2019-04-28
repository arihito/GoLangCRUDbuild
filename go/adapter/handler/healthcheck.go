package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	IHealthCheckHandler interface {
		Get(ctx echo.Context) error
	}

	HealthCheckHandler struct {
	}
)

func NewHealthCheckHandler() IHealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Get(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
