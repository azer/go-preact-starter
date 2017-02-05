package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// Set all your API endpoints in this function, on top of the file
func EnableAPI(server *echo.Echo) {
	server.GET("/api/hi", Hi)
}

// And create the handlers after that
func Hi(c echo.Context) error {
	return c.JSON(http.StatusOK, &struct {
		Hi string
	}{"there"})
}
