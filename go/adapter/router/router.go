package router

import (
	"github.com/labstack/echo"
	"github.com/sminoeee/sample-app/go/adapter/handler"
)

// Routing
// APIバージョンもあるので分けても group に分けて良いかも
func Router(e *echo.Echo) *echo.Echo {
	// prefix: /api
	g := e.Group("/api")

	// health check
	healthCheckHandler := handler.NewHealthCheckHandler()
	g.GET("/healthcheck", healthCheckHandler.Get)

	// users
	userHandler := handler.NewUserHandler()
	g.GET("/users/:id", userHandler.Get)
	g.POST("/users", userHandler.Store)

	// dashboard
	dashboardHandler := handler.NewDashboardHandler()
	g.GET("/dashboard", dashboardHandler.Get)

	// study_groups

	// TODO 実装

	return e
}
