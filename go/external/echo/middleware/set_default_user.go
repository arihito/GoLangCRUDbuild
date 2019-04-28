package middleware

import (
	"github.com/labstack/echo"
	"github.com/sminoeee/sample-app/go/util"
)

// Sample-app 内で使用するユーザー情報を echo.Context に設定する
// デフォルトとして users.id: 1 を request header に詰める
func SetDefaultUser(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defaultUserID := "1" // request header は string のみ設定可能
		c.Request().Header.Set(util.SessionUserID, defaultUserID)
		return h(c)
	}
}
