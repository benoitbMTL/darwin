package main

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/labstack/echo/v4"
)

func registerActions(e *echo.Echo) {

	// COMMAND INJECTION
	e.POST("/command-injection", handleCommandInjectionAction)

	// SQL INJECTION
	e.POST("/sql-injection", handleSQLInjectionAction)

	// BOT DECEPTION
	e.GET("/view-page-source", handleViewPageSourceAction)
	e.GET("/bot-deception", handleBotDeceptionAction)

	// PING
	e.POST("/ping", handlePingAction)
}

///////////////////////////////////////////////////////////////////////////////////
// COMMAND INJECTION                                                             //
///////////////////////////////////////////////////////////////////////////////////

func handleCommandInjectionAction(c echo.Context) error {
	username := c.FormValue("username")

	password, ok := UserPassMap[username]
	if !ok {
		return c.String(http.StatusBadRequest, "Invalid username")
	}

	// Perform Authentication
	cmd := exec.Command("curl", DVWA_URL+"/login.php",
		"-H", "authority: "+DVWA_HOST,
		"-H", "cache-control: max-age=0",
		"-H", "content-type: application/x-www-form-urlencoded",
		"-H", "origin: "+DVWA_URL,
		"-H", "referer: "+DVWA_URL+"/",
		"-H", "user-agent: "+USER_AGENT,
		"--insecure",
		"--silent",
		"--data-raw", "username="+username+"&password="+password+"&Login=Login",
		"-c", "cookie.txt",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Execute Command Injection
	cmd2 := exec.Command("curl", DVWA_URL+"/vulnerabilities/exec/",
		"-H", "authority: "+DVWA_HOST,
		"-H", "cache-control: max-age=0",
		"-H", "content-type: application/x-www-form-urlencoded",
		"-H", "origin: "+DVWA_URL,
		"-H", "referer: "+DVWA_URL+"/index.php",
		"-H", "user-agent: "+USER_AGENT,
		"--insecure",
		"--silent",
		"--data-raw", "ip=;ls&Submit=Submit",
		"-b", "cookie.txt",
	)

	output2, err := cmd2.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the HTML content of the two curl command
	return c.HTML(http.StatusOK, string(output)+"\n"+string(output2))

}

///////////////////////////////////////////////////////////////////////////////////
// SQL INJECTION                                                                 //
///////////////////////////////////////////////////////////////////////////////////

func handleSQLInjectionAction(c echo.Context) error {
	username := c.FormValue("username")

	password, ok := UserPassMap[username]
	if !ok {
		return c.String(http.StatusBadRequest, "Invalid username")
	}

	// Perform Authentication
	cmd := exec.Command("curl", DVWA_URL+"/login.php",
		"-H", "authority: "+DVWA_HOST,
		"-H", "cache-control: max-age=0",
		"-H", "content-type: application/x-www-form-urlencoded",
		"-H", "origin: "+DVWA_URL,
		"-H", "referer: "+DVWA_URL+"/",
		"-H", "user-agent: "+USER_AGENT,
		"--insecure",
		"--silent",
		"--data-raw", "username="+username+"&password="+password+"&Login=Login",
		"-c", "cookie.txt",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Execute SQL Injection
	cmd2 := exec.Command("curl", DVWA_URL+"/vulnerabilities/sqli/?id=%27OR+1%3D1%23&Submit=Submit",
		"-H", "authority: "+DVWA_HOST,
		"-H", "cache-control: max-age=0",
		"-H", "content-type: application/x-www-form-urlencoded",
		"-H", "origin: "+DVWA_URL,
		"-H", "referer: "+DVWA_URL+"/index.php",
		"-H", "user-agent: "+USER_AGENT,
		"--insecure",
		"--silent",
		"-b", "cookie.txt",
	)

	output2, err := cmd2.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the HTML content of the two curl command
	return c.HTML(http.StatusOK, string(output)+"\n"+string(output2))

}

///////////////////////////////////////////////////////////////////////////////////
// BOT DECEPTION                                                                 //
///////////////////////////////////////////////////////////////////////////////////

func handleViewPageSourceAction(c echo.Context) error {
	// Execute curl command to get the source code of login.php
	cmd := exec.Command("curl", "-s", "-k", DVWA_URL+"/login.php",
		"-H", "authority: "+DVWA_HOST,
		"-H", "cache-control: max-age=0",
		"-H", "content-type: application/x-www-form-urlencoded",
		"-H", "origin: "+DVWA_URL,
		"-H", "referer: "+DVWA_URL,
		"-H", "user-agent: FortiWeb Demo Tool",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Split the output into lines and get the last 15 lines
	lines := strings.Split(string(output), "\n")
	lastLines := lines[len(lines)-15:]

	// Return the last 15 lines of the source code
	return c.String(http.StatusOK, strings.Join(lastLines, "\n"))
}

func handleBotDeceptionAction(c echo.Context) error {
	// Execute curl command to get the fake_url.php page
	cmd := exec.Command("curl", "-s", "-k", DVWA_URL+"/fake_url.php",
		"-H", "authority: "+DVWA_HOST,
		"-H", "cache-control: max-age=0",
		"-H", "content-type: application/x-www-form-urlencoded",
		"-H", "origin: "+DVWA_URL,
		"-H", "referer: "+DVWA_URL+"/index.php",
		"-H", "user-agent: FortiWeb Demo Tool",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the HTML content of the fake_url.php page
	return c.HTML(http.StatusOK, string(output))
}

///////////////////////////////////////////////////////////////////////////////////
// PING                                                                          //
///////////////////////////////////////////////////////////////////////////////////

func handlePingAction(c echo.Context) error {
	ipFqdn := c.FormValue("ip-fqdn")

	// Sanitize the input to avoid command injection
	if strings.ContainsAny(ipFqdn, ";&|") {
		return c.String(http.StatusBadRequest, "Invalid characters in input")
	}

	// Execute the ping command
	cmd := exec.Command("ping", "-c", "1", ipFqdn)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the output of the ping command
	return c.String(http.StatusOK, string(output))
}
