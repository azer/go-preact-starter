package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"io/ioutil"
	"os"
)

func init() {
	// Read env variables from .env file in the root directory
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	server := echo.New()

	EnablePreact(server) // Enables server-side rendering for Preact components
	EnableAPI(server)

	server.Logger.SetOutput(ioutil.Discard)

	// Setup static files
	server.Static("/public", "./public")
	server.GET("/favicon.ico", func(c echo.Context) error {
		c.Redirect(301, "/public/favicon.ico")
		return nil
	})

	// Run the server. You can edit the ADDR from .env file on the project folder.
	server.Start(os.Getenv("ADDR"))
}
