package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Name                 string `json:"name,omitempty"`
	Vip                  string `json:"vip,omitempty"`
	Status               string `json:"status,omitempty"`
	Interface            string `json:"interface,omitempty"`
	ServerBalance        string `json:"server-balance,omitempty"`
	Health               string `json:"health,omitempty"`
	IP                   string `json:"ip,omitempty"`
	SSL                  string `json:"ssl,omitempty"`
	Port                 int    `json:"port,omitempty"`
	SignatureRule        string `json:"signature-rule,omitempty"`
	XForwardedForRule    string `json:"x-forwarded-for-rule,omitempty"`
	DeploymentMode       string `json:"deployment-mode,omitempty"`
	Protocol             string `json:"protocol,omitempty"`
	ImplicitSSL          string `json:"implicit_ssl,omitempty"`
	VServer              string `json:"vserver,omitempty"`
	Service              string `json:"service,omitempty"`
	WebProtectionProfile string `json:"web-protection-profile,omitempty"`
	ServerPool           string `json:"server-pool,omitempty"`
	Tlog                 string `json:"tlog,omitempty"`
	HttpsService         string `json:"https-service,omitempty"`
	Certificate          string `json:"certificate,omitempty"`
	XForwardedForSupport string `json:"x-forwarded-for-support,omitempty"`
	MatchObject          string `json:"match-object,omitempty"`
	MatchCondition       string `json:"match-condition,omitempty"`
	MatchExpression      string `json:"match-expression,omitempty"`
	Concatenate          string `json:"concatenate,omitempty"`
	ProfileInherit       string `json:"profile-inherit,omitempty"`
	IsDefault            string `json:"is-default,omitempty"`
	Action               string `json:"action,omitempty"`
	Location             string `json:"location,omitempty"`
	Object               string `json:"object,omitempty"`
	RegExp               string `json:"regexp,omitempty"`
}

type Request struct {
	Data Data `json:"data"`
}

func createVirtualIP(host, token, vipName, vip, iface string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/system/vip", host)

	data := Data{
		Name:      vipName,
		Vip:       vip,
		Interface: iface,
	}

	log.Printf("Creating Virtual IP: %s\n", vipName)
	return sendRequest("POST", url, token, data)
}

func createNewServerPool(host, token, poolName string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool", host)

	data := Data{
		Name:          poolName,
		ServerBalance: "enable",
		Health:        "HLTHCK_HTTP",
	}

	sendRequest("POST", url, token, data)
}

func createNewMemberPool(host, token, poolName, ip string, port int) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool/pserver-list?mkey=%s", host, poolName)

	data := Data{
		IP:   ip,
		SSL:  "enable",
		Port: port,
	}

	sendRequest("POST", url, token, data)
}

func createNewVirtualServer(host, token, vServerName string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/vserver", host)

	data := Data{
		Name: vServerName,
	}

	sendRequest("POST", url, token, data)
}

func assignVIPtoVirtualServer(host, token, vServerName, iface, vipStatus, vipName string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/vserver/vip-list?mkey=%s", host, vServerName)

	data := Data{
		Interface: iface,
		Status:    vipStatus,
		Vip:       vipName,
	}

	sendRequest("POST", url, token, data)
}

func cloneSignatureStandardProtection(host, token, mkey, cloneMkey string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/signature?mkey=%s&clone_mkey=%s", host, mkey, cloneMkey)
	sendRequest("POST", url, token, Data{})
}

func createNewXForwardedForRule(host, token, ruleName string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/x-forwarded-for", host)

	data := Data{
		Name:                 ruleName,
		XForwardedForSupport: "enable",
	}

	sendRequest("POST", url, token, data)
}

func cloneInlineProtectionProfile(host, token, mkey, cloneMkey string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/web-protection-profile.inline-protection?mkey=%s&clone_mkey=%s", host, url.QueryEscape(mkey), url.QueryEscape(cloneMkey))

	sendRequest("POST", url, token, Data{})
}

func configureProtectionProfile(host, token, signatureRule, xForwardedForRule string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/web-protection-profile.inline-protection?mkey=STANDARD_PROTECTION_CLONE", host)

	data := Data{
		SignatureRule:     signatureRule,
		XForwardedForRule: xForwardedForRule,
	}

	sendRequest("PUT", url, token, data)
}

func createNewPolicy(host, token, name, deploymentMode, protocol, ssl, implicitSsl, vserver, service, webProtectionProfile, serverPool, tlog, httpsService, certificate string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/policy", host)

	data := Data{
		Name:                 name,
		DeploymentMode:       deploymentMode,
		Protocol:             protocol,
		SSL:                  ssl,
		ImplicitSSL:          implicitSsl,
		VServer:              vserver,
		Service:              service,
		WebProtectionProfile: webProtectionProfile,
		ServerPool:           serverPool,
		Tlog:                 tlog,
		HttpsService:         httpsService,
		Certificate:          certificate,
	}

	sendRequest("POST", url, token, data)
}

func createNewHTTPContentRoute(host, token, routeName, serverPool string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/http-content-routing-policy", host)

	data := Data{
		Name:       routeName,
		ServerPool: serverPool,
	}

	sendRequest("POST", url, token, data)
}

func createNewContentRouteMatchingCriteria(host, token, routeName, matchObject, matchCondition, matchExpression, concatenate string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/http-content-routing-policy/content-routing-match-list?mkey=%s", host, routeName)

	data := Data{
		MatchObject:     matchObject,
		MatchCondition:  matchCondition,
		MatchExpression: matchExpression,
		Concatenate:     concatenate,
	}

	sendRequest("POST", url, token, data)
}

func assignContentRouteToPolicy(host, token, policyName, routeName, profileInherit, webProtectionProfile, isDefault, status string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/policy/http-content-routing-list?mkey=%s", host, policyName)

	data := Data{
		Name:                 routeName,
		ProfileInherit:       profileInherit,
		WebProtectionProfile: webProtectionProfile,
		IsDefault:            isDefault,
		Status:               status,
	}

	sendRequest("POST", url, token, data)
}

func createNewURLRewriteRule(host, token, ruleName, action, location string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/url-rewrite.url-rewrite-rule", host)

	data := Data{
		Name:     ruleName,
		Action:   action,
		Location: location,
	}

	sendRequest("POST", url, token, data)
}

func createNewMatchCondition(host, token, ruleName, object, regExp string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/url-rewrite.url-rewrite-rule/match-condition?mkey=%s", host, ruleName)

	data := Data{
		Object: object,
		RegExp: regExp,
	}

	sendRequest("POST", url, token, data)
}

func createNewRewritePolicy(host, token, policyName string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/url-rewrite.url-rewrite-policy", host)

	data := Data{
		Name: policyName,
	}

	sendRequest("POST", url, token, data)
}

func createNewRewritePolicyRule(host, token, policyName, ruleName string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/url-rewrite.url-rewrite-policy/rule?mkey=%s", host, policyName)

	data := Data{
		Name: ruleName,
	}

	sendRequest("POST", url, token, data)
}

func addRewritingRuleToProtectionProfile(host, token, profileName, policyName string) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/web-protection-profile.inline-protection?mkey=%s", host, profileName)

	data := Data{
		Name: policyName,
	}

	sendRequest("PUT", url, token, data)
}

func sendRequest(method, url, token string, data Data) ([]byte, error) {
	log.Printf("Method: %s\n", method)
    log.Printf("URL: %s\n", url)
    log.Printf("Token: %s\n", token)
    log.Printf("Data: %+v\n", data)

	reqData := Request{
		Data: data,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		log.Printf("Error marshalling request data: %v\n", err)
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return nil, err
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/json")

	// Create a custom HTTP client with SSL/TLS certificate verification disabled
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	log.Printf("Sending %s request to: %s\n", method, url)

	time.Sleep(time.Duration(500) * time.Millisecond)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	log.Printf("Response received: %s\n", string(body))

	return body, nil
}

func calculateToken() string {
	tokenData := fmt.Sprintf(`{"username":"%s","password":"%s","vdom":"%s"}`, USERNAME_API, PASSWORD_API, VDOM_API)
	return base64.StdEncoding.EncodeToString([]byte(tokenData))
}

func onboardNewApplicationPolicy(c echo.Context) error {
	host := FWB_MGT_IP
	token := calculateToken()
	log.Printf("Token: %s\n", token)

	data := Data{
		Name:      "VIP1",
		Vip:       "192.168.4.80/24",
		Interface: "port1",
	}
	log.Printf("Data: %+v\n", data)  // Add this line

	result, err := createVirtualIP(host, token, data)
	if err != nil {
		// Handle the error
		log.Printf("Error creating virtual IP: %v\n", err)
		return err
	}

	log.Printf("New Application Policy onboarded successfully\n")
	return c.JSON(http.StatusOK, string(result))
}
