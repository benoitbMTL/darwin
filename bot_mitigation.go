package main

import (
	"crypto/tls"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

///////////////////////////////////////////////////////////////////////////////////
// BOT DECEPTION                                                                 //
///////////////////////////////////////////////////////////////////////////////////

func handleViewPageSourceAction(c echo.Context) error {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	req, err := http.NewRequest("GET", DVWA_URL+"/login.php", nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	req.Header.Set("authority", DVWA_HOST)
	req.Header.Set("origin", DVWA_URL)
	req.Header.Set("referer", DVWA_URL)
	req.Header.Set("user-agent", USER_AGENT)
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	lines := strings.Split(string(output), "\n")
	lastLines := lines[len(lines)-15:]
	lastLinesString := strings.Join(lastLines, "\n")

	// Check if specific strings exist, if so wrap them with HTML tags for red color
	if strings.Contains(lastLinesString, `href="/fake_url.php"`) {
		lastLinesString = strings.ReplaceAll(lastLinesString, `href="/fake_url.php"`, `<span style="color:red;">href="/fake_url.php"</span>`)
	}

	if strings.Contains(lastLinesString, `style='display:none'`) {
		lastLinesString = strings.ReplaceAll(lastLinesString, `style='display:none'`, `<span style="color:red;">style='display:none'</span>`)
	}

	// Wrap the result with pre tags for courier new and font size
	result := fmt.Sprintf(`<pre style="font-family:'Courier New', monospace; font-size:14px;">%s</pre>`, html.EscapeString(lastLinesString))

	return c.HTML(http.StatusOK, result)
}

func handleBotDeceptionAction(c echo.Context) error {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	req, err := http.NewRequest("GET", DVWA_URL+"/fake_url.php", nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	req.Header.Set("authority", DVWA_HOST)
	req.Header.Set("origin", DVWA_URL)
	req.Header.Set("referer", DVWA_URL+"/index.php")
	req.Header.Set("user-agent", USER_AGENT)
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, string(output))
}
