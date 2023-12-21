package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Define all your variables here
var (
	// FortiWeb Tool Configuration
	DVWA_URL      = "https://dvwa.corp.fabriclab.ca"
	DVWA_HOST     = "dvwa.corp.fabriclab.ca"
	JUICESHOP_URL = "https://shop.corp.fabriclab.ca"
	FWB_URL       = "https://fwb.corp.fabriclab.ca/fwb/"
	SPEEDTEST_URL = "https://speedtest.corp.fabriclab.ca"
	PETSTORE_URL  = "https://petstore.corp.fabriclab.ca/api/v3/pet"
	USERNAME_API  = "userapi"
	PASSWORD_API  = "userAPI123!"
	VDOM_API      = "root"
	FWB_MGT_IP    = "10.163.7.21"
	POLICY        = "DVWA_POLICY"
	USER_AGENT    = "FortiWeb Demo Tool"

	// DVWA Credentials
	UserPassMap = map[string]string{
		"admin":   "password",
		"gordonb": "abc123",
		"1337":    "charley",
		"pablo":   "letmein",
		"smithy":  "password",
	}

	// Credentials for Credential Stuffing Demonstration
	CredentialStuffingMap = map[string]string{
		"pklangdon4@msn.com":             "ppl11266",
		"muldersstan@gmail.com":          "renzo1205",
		"forsternp2@aol.com":             "freedom1",
		"cragsy@msn.com":                 "Snatch01",
		"bjrehdorf@hotmail.com":          "Apollo25504",
		"baz2709@icloud.com":             "sophie12",
		"amysiura@ymail.com":             "active95",
		"jond714@gmail.com":              "jonloveslax",
		"josefahorenstein87@hotmail.com": "qel737Xf3",
		"bizotic6@gmail.com":             "snaker26",
	}

	// Coutries for Web Scanning, FortiView Country and GeoLocation Demonstrations
	ipCountryMap = map[string]string{
		"Argentina":      "103.50.33.74",
		"Australia":      "103.137.12.51",
		"Brazil":         "193.19.205.129",
		"Canada":         "45.88.190.128",
		"Chile":          "85.190.229.4",
		"France":         "89.40.183.21",
		"Germany":        "5.180.62.6",
		"Italy":          "138.199.54.247",
		"Japan":          "156.146.35.84",
		"Mexico":         "185.153.177.108",
		"Norway":         "51.13.51.13",
		"Spain":          "31.13.188.139",
		"Ukraine":        "37.19.218.159",
		"United Kingdom": "5.101.138.227",
		"United States":  "45.14.195.100",
	}

	// REST API Demo
	VipName                             = "VIRTUAL_IP"
	VipIp                               = "10.163.7.60/24"
	Interface                           = "port1"
	PoolName                            = "SERVER_POOL"
	ServerBalance                       = "enable"
	HealthCheck                         = "HLTHCK_HTTP"
	PoolMemberIPs                       = []string{"10.0.0.10", "10.0.0.20", "10.0.0.30"}
	PoolMemberSSL                       = "disable"
	PoolMemberPort                      = 4000
	VirtualServerName                   = "VIRTUAL_SERVER"
	VipStatus                           = "enable"
	XForwardedForName                   = "X_FORWARDED_FOR_RULE"
	XForwardedForSupport                = "enable"
	OriginalSignatureProtectionName     = "Standard Protection"
	CloneSignatureProtectionName        = "STANDARD_PROTECTION_CLONE"
	OriginalInlineProtectionProfileName = "Inline Standard Protection"
	CloneInlineProtectionProfileName    = "INLINE_STANDARD_PROTECTION_CLONE"
	PolicyName                          = "NEW_POLICY"
	PolicyDeploymentMode                = "server-pool"
	PolicyProtocol                      = "HTTP"
	PolicySSL                           = "enable"
	PolicyImplicitSSL                   = "enable"
	PolicyVirtualServer                 = VirtualServerName
	PolicyService                       = "HTTP"
	PolicyInlineProtectionProfile       = CloneInlineProtectionProfileName
	PolicyServerPool                    = PoolName
	PolicyTrafficLog                    = "enable"
	PolicyHTTPSService                  = "HTTPS"
	//PolicyCertificate                   = "newapp.fabriclab.ca"
)

type Config struct {
	DVWA_URL      string `json:"dvwa_url"`
	DVWA_HOST     string `json:"dvwa_host"`
	JUICESHOP_URL string `json:"juiceshop_url"`
	FWB_URL       string `json:"fwb_url"`
	SPEEDTEST_URL string `json:"speedtest_url"`
	PETSTORE_URL  string `json:"petstore_url"`
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
	// fmt.Println(token) // Print the token to the console

	return Config{
		DVWA_URL:      DVWA_URL,
		DVWA_HOST:     DVWA_HOST,
		JUICESHOP_URL: JUICESHOP_URL,
		FWB_URL:       FWB_URL,
		SPEEDTEST_URL: SPEEDTEST_URL,
		PETSTORE_URL:  PETSTORE_URL,
		USERNAME_API:  USERNAME_API,
		PASSWORD_API:  PASSWORD_API,
		VDOM_API:      VDOM_API,
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
	JUICESHOP_URL = currentConfig.JUICESHOP_URL
	FWB_URL = currentConfig.FWB_URL
	SPEEDTEST_URL = currentConfig.SPEEDTEST_URL
	PETSTORE_URL = currentConfig.PETSTORE_URL
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
	JUICESHOP_URL = newConfig.JUICESHOP_URL
	FWB_URL = newConfig.FWB_URL
	SPEEDTEST_URL = newConfig.SPEEDTEST_URL
	PETSTORE_URL = newConfig.PETSTORE_URL
	USERNAME_API = newConfig.USERNAME_API
	PASSWORD_API = newConfig.PASSWORD_API
	VDOM_API = newConfig.VDOM_API
	FWB_MGT_IP = newConfig.FWB_MGT_IP
	POLICY = newConfig.POLICY
	USER_AGENT = newConfig.USER_AGENT

	// Recalculate the TOKEN
	tokenData := fmt.Sprintf(`{"username":"%s","password":"%s","vdom":"%s"}`, USERNAME_API, PASSWORD_API, VDOM_API)
	currentConfig.TOKEN = base64.StdEncoding.EncodeToString([]byte(tokenData))
    // fmt.Println(currentConfig.TOKEN) // Print the encoded token to the console

	return c.JSON(http.StatusOK, currentConfig) // Return currentConfig
}
