package main

import (
	"crypto/tls"
	"io"
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

	output, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	lines := strings.Split(string(output), "\n")
	lastLines := lines[len(lines)-15:]

	return c.String(http.StatusOK, strings.Join(lastLines, "\n"))
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

	output, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, string(output))
}
