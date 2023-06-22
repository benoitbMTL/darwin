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

	////////////////////////////////////////////////////////////////////////
	// ACTION #1 PING                                                     //
	////////////////////////////////////////////////////////////////////////

	// Route to handle the ping form submission
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

	////////////////////////////////////////////////////////////////////////
	// ACTION #2 COMMAND INJECTION                                        //
	////////////////////////////////////////////////////////////////////////

	// Route to handle the command injection action
	e.POST("/command-injection", func(c echo.Context) error {
		username := c.FormValue("username")

		// Map usernames to passwords
		userPassMap := map[string]string{
			"admin":   "password",
			"gordonb": "abc123",
			"1337":    "charley",
			"pablo":   "letmein",
			"smithy":  "password",
		}

		password, ok := userPassMap[username]
		if !ok {
			return c.String(http.StatusBadRequest, "Invalid username")
		}

		// Execute the curl command
		cmd := exec.Command("curl", "https://192.168.4.10/login.php",
			"-H", "authority: 192.168.4.10",
			"-H", "cache-control: max-age=0",
			"-H", "content-type: application/x-www-form-urlencoded",
			"-H", "origin: https://192.168.4.10",
			"-H", "referer: https://192.168.4.10/",
			"-H", "user-agent: FortiWeb Demo Script",
			"--insecure",
			"--data-raw", "username="+username+"&password="+password+"&Login=Login",
			"-c", "cookie.txt",
		)

		output, err := cmd.CombinedOutput()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Return the output of the curl command
		return c.String(http.StatusOK, string(output))
	})

	////////////////////////////////////////////////////////////////////////
	// START THE SERVER                                                   //
	////////////////////////////////////////////////////////////////////////

	// Start the server
	e.Start(":8080")
}
