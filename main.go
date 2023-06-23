package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Set the location of the templates directory
	e.Static("/templates", "templates")

	// Serve static files from the static directory
	e.Static("/static", "static")

	// Route to render the HTML page
	e.GET("/", func(c echo.Context) error {
		return c.File("templates/index.html")
	})

	// Register actions
	registerActions(e)
	
	// Register config handler
	e.GET("/api/config", ConfigHandler)

	// Start the server
	e.Start(":8080")
}
