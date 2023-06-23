package main

var (
	DVWA_URL      = "https://192.168.4.10"
	DVWA_HOST     = "192.168.4.10"
	SHOP_URL      = "https://shop.corp.fabriclab.ca"
	FWB_URL       = "https://192.168.4.10/fwb/"
	SPEEDTEST_URL = "http://speedtest.corp.fabriclab.ca"
	KALI_URL      = "https://flbr1kali01.fortiweb.fabriclab.ca"
	TOKEN         = "eyJ1c2VybmFtZSI6InVzZXJhcGkiLCJwYXNzd29yZCI6ImZhY2VMT0NLeWFybjY3ISJ9Cg=="
	FWB_MGT_IP    = "192.168.4.2"
	POLICY        = "DVWA_POLICY"
	USER_AGENT    = "FortiWeb Demo Tool"

	// Map usernames to passwords
	UserPassMap = map[string]string{
		"admin":   "password",
		"gordonb": "abc123",
		"1337":    "charley",
		"pablo":   "letmein",
		"smithy":  "password",
	}
)
