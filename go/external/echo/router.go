package echo

import (
	"github.com/labstack/echo"
	"github.com/sminoeee/sample-app/go/adapter/handler"
	"github.com/sminoeee/sample-app/go/external/db"
	"github.com/sminoeee/sample-app/go/usecase"
	"net/http"
)

// Routing
// APIバージョンもあるので分けても group に分けて良いかも
func router(e *echo.Echo) *echo.Echo {
	// sample handler
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// seminars
	seminarHandler := handler.NewSeminarHandler(usecase.NewSeminarUseCase(db.Conn))
	e.GET("/seminars/:id", seminarHandler.FindById)
	e.POST("/seminars", seminarHandler.Store)

	// users
	userHandler := handler.NewUserHandler(usecase.NewUserUseCase(db.Conn))
	e.GET("/users/:id", userHandler.FindById)
	e.POST("/users", userHandler.Store)

	// ...

	return e
}
