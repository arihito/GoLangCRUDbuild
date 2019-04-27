package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sminoeee/sample-app/go/adapter/handler"
	"net/http"
	"time"
)

type ApiError struct {
	StatusCode  int    `json:"status_code"`
	Message     string `json:"message"`
	RespondedAt string `json:"responded_at"` // RFC3339
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := ""

	switch e := err.(type) {
	case *handler.ApplicationError:
		// business error
		code = e.Code
		message = e.Message
	default:
		// unknown error
		log.Error(err)
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = http.StatusText(he.Code)
		} else {
			// default code: 500
			message = http.StatusText(code)
		}
	}

	body := ApiError{
		StatusCode:  code,
		Message:     message,
		RespondedAt: time.Now().Format(time.RFC3339),
	}

	c.JSON(code, body)
}
