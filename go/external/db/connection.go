package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sminoeee/sample-app/go/util"
)

var Conn *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	// NOTE: docker 外で動かすときは db -> 127.0.0.1 に変更する
	host := "127.0.0.1"
	if util.IsOnDocker() {
		host = "db"
	}
	Conn, err = gorm.Open("mysql", "sample_user:12345@tcp(" +host+":3306)/sample_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	// デバッグログ: ON
	Conn.LogMode(true)
	return Conn
}
