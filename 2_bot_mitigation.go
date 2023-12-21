package main

import (
	"crypto/tls"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

///////////////////////////////////////////////////////////////////////////////////
// BOT THRESHOLDS                                                                //
///////////////////////////////////////////////////////////////////////////////////

// randomPath generates a random path of a specified length
func randomPath(minLength, maxLength int) string {
    var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789-_+?=")
    var startingLetters = []rune("abcdefghijklmnopqrstuvwxyz")

    source := rand.NewSource(time.Now().UnixNano())
    random := rand.New(source)

    length := random.Intn(maxLength-minLength+1) + minLength

    var sb strings.Builder
    sb.WriteRune(startingLetters[random.Intn(len(startingLetters))]) // Start with a-z

    for i := 1; i < length; i++ {
        sb.WriteRune(letters[random.Intn(len(letters))])
    }

    return sb.String()
}

func handleBotThresholdAction(c echo.Context) error {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	// Loop to send 150 GET requests with random paths
	for i := 0; i < 150; i++ {
		randomPath := randomPath(5, 10) // Generate a random path
		url := DVWA_URL + "/" + randomPath

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Set headers to mimic a regular HTTP request
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		req.Header.Set("Accept-Language", "en-US,en;q=0.5")
		req.Header.Set("X-Forwarded-For", "212.64.64.64") // Set IP to simulate request from China

		resp, err := client.Do(req)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer resp.Body.Close()

		_, err = io.ReadAll(resp.Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Optional: process the response or log it
	}

	// Return a successful HTML response
	return c.HTML(http.StatusOK, "<p>Requests completed.</p>")
}



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
