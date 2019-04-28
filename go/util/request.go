package util

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"strconv"
)

const SessionUserID = "X-SESSION-USER-ID"

// echo.Context から ユーザー を取得する
func GetUserIDFromRequest(c echo.Context) int64 {
	userID, err := strconv.Atoi(c.Request().Header.Get(SessionUserID))
	if err != nil {
		log.Error("UserID missing.")
		return 1 // ヘッダーがから取れなければデフォルト1
	}
	return int64(userID)
}
