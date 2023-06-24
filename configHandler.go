package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var UserPassMap = map[string]string{
	"admin":   "password",
	"gordonb": "abc123",
	"1337":    "charley",
	"pablo":   "letmein",
	"smithy":  "password",
}

type Config struct {
	DVWA_URL      string `json:"dvwa_url"`
	DVWA_HOST     string `json:"dvwa_host"`
	SHOP_URL      string `json:"shop_url"`
	FWB_URL       string `json:"fwb_url"`
	SPEEDTEST_URL string `json:"speedtest_url"`
	KALI_URL      string `json:"kali_url"`
	TOKEN         string `json:"token"`
	FWB_MGT_IP    string `json:"fwb_mgt_ip"`
	POLICY        string `json:"policy"`
	USER_AGENT    string `json:"user_agent"`
}

var defaultConfig = Config{
	DVWA_URL:      "https://192.168.4.10",
	DVWA_HOST:     "192.168.4.10",
	SHOP_URL:      "https://shop.corp.fabriclab.ca",
	FWB_URL:       "https://192.168.4.10/fwb/",
	SPEEDTEST_URL: "http://speedtest.corp.fabriclab.ca",
	KALI_URL:      "https://flbr1kali01.fortiweb.fabriclab.ca",
	TOKEN:         "eyJ1c2VybmFtZSI6InVzZXJhcGkiLCJwYXNzd29yZCI6ImZhY2VMT0NLeWFybjY3ISJ9Cg==",
	FWB_MGT_IP:    "192.168.4.2",
	POLICY:        "DVWA_POLICY",
	USER_AGENT:    "FortiWeb Demo Tool",
}

func ConfigHandler(c echo.Context) error {
	// Here you would normally retrieve the current configuration from where you have it stored
	// For this example, we'll just return the default configuration
	return c.JSON(http.StatusOK, defaultConfig)
}

func DefaultConfigHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, defaultConfig)
}

func SaveConfigHandler(c echo.Context) error {
	// Parse the request body into a Config struct
	var newConfig Config
	if err := c.Bind(&newConfig); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Save the new configuration values
	defaultConfig = newConfig

	// Return a success response
	return c.JSON(http.StatusOK, newConfig)
}
