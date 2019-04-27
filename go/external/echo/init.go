package echo

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	middleware2 "github.com/sminoeee/sample-app/go/external/echo/middleware"
	"os"
)

func Init() {

	e := echo.New()

	// logging
	e.Debug = true
	e.Logger.SetOutput(os.Stdout) // 一旦標準出力で

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// routing
	e = router(e)

	// http error handler
	e.HTTPErrorHandler = middleware2.CustomHTTPErrorHandler

	// ### start ###
	e.Logger.Info(e.Start(":1323"))
}