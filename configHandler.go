package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Define all your variables here
var (
	DVWA_URL      = "https://192.168.4.10"
	DVWA_HOST     = "192.168.4.10"
	SHOP_URL      = "https://shop.corp.fabriclab.ca"
	FWB_URL       = "https://192.168.4.10/fwb/"
	SPEEDTEST_URL = "http://speedtest.corp.fabriclab.ca"
	KALI_URL      = "https://flbr1kali01.fortiweb.fabriclab.ca"
	USERNAME_API  = "userapi"
	PASSWORD_API  = "abc123"
	VDOM_API      = "root"
	FWB_MGT_IP    = "192.168.4.2"
	POLICY        = "DVWA_POLICY"
	USER_AGENT    = "FortiWeb Demo Tool"

	UserPassMap = map[string]string{
		"admin":   "password",
		"gordonb": "abc123",
		"1337":    "charley",
		"pablo":   "letmein",
		"smithy":  "password",
	}
)

type Config struct {
	DVWA_URL      string `json:"dvwa_url"`
	DVWA_HOST     string `json:"dvwa_host"`
	SHOP_URL      string `json:"shop_url"`
	FWB_URL       string `json:"fwb_url"`
	SPEEDTEST_URL string `json:"speedtest_url"`
	KALI_URL      string `json:"kali_url"`
	USERNAME_API  string `json:"username_api"`
	PASSWORD_API  string `json:"password_api"`
	VDOM_API      string `json:"vdom_api"`
	TOKEN         string `json:"token"`
	FWB_MGT_IP    string `json:"fwb_mgt_ip"`
	POLICY        string `json:"policy"`
	USER_AGENT    string `json:"user_agent"`
}

func initialConfig() Config {
	tokenData := fmt.Sprintf(`{"username":"%s","password":"%s","vdom":"%s"}`, USERNAME_API, PASSWORD_API, VDOM_API)
	token := base64.StdEncoding.EncodeToString([]byte(tokenData))

	return Config{
		DVWA_URL:      DVWA_URL,
		DVWA_HOST:     DVWA_HOST,
		SHOP_URL:      SHOP_URL,
		FWB_URL:       FWB_URL,
		SPEEDTEST_URL: SPEEDTEST_URL,
		KALI_URL:      KALI_URL,
		TOKEN:         token,
		FWB_MGT_IP:    FWB_MGT_IP,
		POLICY:        POLICY,
		USER_AGENT:    USER_AGENT,
	}
}

var defaultConfig = initialConfig()
var currentConfig = initialConfig()

func ConfigHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, currentConfig) // Use currentConfig
}

func DefaultConfigHandler(c echo.Context) error {
	currentConfig = defaultConfig // Reset currentConfig
	DVWA_URL = currentConfig.DVWA_URL
	DVWA_HOST = currentConfig.DVWA_HOST
	SHOP_URL = currentConfig.SHOP_URL
	FWB_URL = currentConfig.FWB_URL
	SPEEDTEST_URL = currentConfig.SPEEDTEST_URL
	KALI_URL = currentConfig.KALI_URL
	USERNAME_API = currentConfig.USERNAME_API
	PASSWORD_API = currentConfig.PASSWORD_API
	VDOM_API = currentConfig.VDOM_API
	FWB_MGT_IP = currentConfig.FWB_MGT_IP
	POLICY = currentConfig.POLICY
	USER_AGENT = currentConfig.USER_AGENT
	return c.JSON(http.StatusOK, currentConfig) // Return currentConfig
}

func SaveConfigHandler(c echo.Context) error {
	var newConfig Config
	if err := c.Bind(&newConfig); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	currentConfig = newConfig // Save changes to currentConfig
	DVWA_URL = newConfig.DVWA_URL
	DVWA_HOST = newConfig.DVWA_HOST
	SHOP_URL = newConfig.SHOP_URL
	FWB_URL = newConfig.FWB_URL
	SPEEDTEST_URL = newConfig.SPEEDTEST_URL
	KALI_URL = newConfig.KALI_URL
	USERNAME_API = newConfig.USERNAME_API
	PASSWORD_API = newConfig.PASSWORD_API
	VDOM_API = newConfig.VDOM_API
	FWB_MGT_IP = newConfig.FWB_MGT_IP
	POLICY = newConfig.POLICY
	USER_AGENT = newConfig.USER_AGENT

// Recalculate the TOKEN
	tokenData := fmt.Sprintf(`{"username":"%s","password":"%s","vdom":"%s"}`, USERNAME_API, PASSWORD_API, VDOM_API)
	currentConfig.TOKEN = base64.StdEncoding.EncodeToString([]byte(tokenData))

	return c.JSON(http.StatusOK, currentConfig) // Return currentConfig
}
