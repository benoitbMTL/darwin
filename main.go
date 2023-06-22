package main

import (
	"net/http"
	"os/exec"
	"strings"

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

	// Route to handle the form submission
	e.POST("/ping", func(c echo.Context) error {
		ipFqdn := c.FormValue("ip-fqdn")

		// Sanitize the input to avoid command injection
		if strings.ContainsAny(ipFqdn, ";&|") {
			return c.String(http.StatusBadRequest, "Invalid characters in input")
		}

		// Execute the ping command
		cmd := exec.Command("ping", "-c", "4", ipFqdn)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Return the output of the ping command
		return c.String(http.StatusOK, string(output))
	})

	// Start the server
	e.Start(":8080")
}
