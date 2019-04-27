package main

import (
	"github.com/sminoeee/sample-app/go/external/db"
	"github.com/sminoeee/sample-app/go/external/echo"
)

func main() {
	// db-connection
	db.ConnectDB()

	// echo, routing, middleware
	echo.Init()

}
