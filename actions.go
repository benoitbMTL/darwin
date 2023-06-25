package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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

	// HEALTH CHECK
	e.GET("/health-check", handleHealthCheckAction)

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
		log.Println("Invalid username") // Log the error
		return c.String(http.StatusBadRequest, "Invalid username")
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	client := &http.Client{
		Jar: jar,
	}

	// Perform Authentication
	data := url.Values{
		"username": {username},
		"password": {password},
		"Login":    {"Login"},
	}
	req, err := http.NewRequest("POST", DVWA_URL+"/login.php", strings.NewReader(data.Encode()))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	req.Header.Set("authority", DVWA_HOST)
	req.Header.Set("origin", DVWA_URL)
	req.Header.Set("referer", DVWA_URL+"/")
	req.Header.Set("user-agent", USER_AGENT)
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		return c.HTML(http.StatusOK, `<pre style="color: red; font-family: 'Courier New', monospace; white-space: pre-wrap;">The Virtual Server is not reachable</pre>`)
	}

	defer resp.Body.Close()

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Log the response body
	log.Print(string(output))

	// Execute Command Injection
	data = url.Values{
		"ip":     {";ls"},
		"Submit": {"Submit"},
	}
	req, err = http.NewRequest("POST", DVWA_URL+"/vulnerabilities/exec/", strings.NewReader(data.Encode()))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	req.Header.Set("authority", DVWA_HOST)
	req.Header.Set("origin", DVWA_URL)
	req.Header.Set("referer", DVWA_URL+"/index.php")
	req.Header.Set("user-agent", USER_AGENT)
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err = client.Do(req)
	if err != nil {
		return c.HTML(http.StatusOK, `<pre style="color: red; font-family: 'Courier New', monospace; white-space: pre-wrap;">The Virtual Server is not reachable</pre>`)
	}

	defer resp.Body.Close()

	output2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Log the response body
	log.Print(string(output2))

	// Return the HTML content
	return c.HTML(http.StatusOK, string(output2))
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
// HEALTH CHECK                                                                  //
///////////////////////////////////////////////////////////////////////////////////

func handleHealthCheckAction(c echo.Context) error {
	urls := []string{DVWA_URL, SHOP_URL, FWB_URL, SPEEDTEST_URL, KALI_URL, "https://www.fortinet.com"}
	result := ""

	// Define a custom HTTP client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Loop over the URLs
	for _, url := range urls {
		res, err := client.Get(url)
		if err != nil {
			// log.Println(fmt.Sprintf("%s is not reachable. Error: %s", url, err.Error())) // Log debug
			result += fmt.Sprintf("<p style=\"color:red\">%s is not reachable. Error: %s</p>", url, err.Error())
		} else {
			// log.Println(fmt.Sprintf("%s is reachable. HTTP Code: %d", url, res.StatusCode)) // Log debug
			result += fmt.Sprintf("<p>%s is reachable. HTTP Code: %d</p>", url, res.StatusCode)
		}
	}

	// Handle FWB_MGT_IP separately because it's only an IP without a scheme
	ip := "http://" + FWB_MGT_IP
	res, err := client.Get(ip)
	if err != nil {
		// log.Println(fmt.Sprintf("%s is not reachable. Error: %s", ip, err.Error())) // Log debug
		result += fmt.Sprintf("<p style=\"color:red\">%s is not reachable. Error: %s</p>", ip, err.Error())
	} else {
		// log.Println(fmt.Sprintf("%s is reachable. HTTP Code: %d", ip, res.StatusCode)) // Log debug
		result += fmt.Sprintf("<p>%s is reachable. HTTP Code: %d</p>", ip, res.StatusCode)
	}

	return c.HTML(http.StatusOK, result)
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
	cmd := exec.Command("ping", "-c", "2", ipFqdn)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Check the error type
		if exitError, ok := err.(*exec.ExitError); ok {
			// Check the exit code
			if exitError.ExitCode() == 1 {
				return c.String(http.StatusInternalServerError, "The destination is not reachable")
			} else if exitError.ExitCode() == 2 {
				return c.String(http.StatusInternalServerError, "The FQDN does not resolve")
			}
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the output of the ping command
	return c.String(http.StatusOK, string(output))
}
