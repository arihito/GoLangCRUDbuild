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
	g.GET("/users/:id/study_groups", userHandler.GetStudyGroups)
	g.POST("/users", userHandler.Store)

	// dashboard
	dashboardHandler := handler.NewDashboardHandler()
	g.GET("/dashboard", dashboardHandler.Get)

	// study_groups
	studyGroupHandler := handler.NewStudyGroupHandler()
	g.POST("/study_groups", studyGroupHandler.Create)       // グループ作成
	g.GET("/study_groups/:id", studyGroupHandler.Get)       // グループ詳細取得
	g.PUT("/study_groups/:id", studyGroupHandler.Update)    // グループ更新
	g.DELETE("/study_groups/:id", studyGroupHandler.Delete) // グループ削除

	// TODO 実装 ...

	// CRUD events
	//

	return e
}
