package main

import (
	"github.com/labstack/echo"
	"os"
	"preact"
)

func EnablePreact(server *echo.Echo) {
	server.Use(preact.HTTPHandler)

	if len(os.Getenv("DEVELOP")) > 0 {
		go preact.WatchCodeChanges()
	}
}
