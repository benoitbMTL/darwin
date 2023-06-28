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

type VirtualIPData struct {
	Name      string `json:"name,omitempty"`
	Vip       string `json:"vip,omitempty"`
	Interface string `json:"interface,omitempty"`
}

type ServerPoolData struct {
	Name          string `json:"name,omitempty"`
	ServerBalance string `json:"server-balance,omitempty"`
	Health        string `json:"health,omitempty"`
}

type MemberPoolData struct {
	IP   string `json:"ip,omitempty"`
	SSL  string `json:"ssl,omitempty"`
	Port int    `json:"port,omitempty"`
}

type Request struct {
	Data interface{} `json:"data"`
}

func createVirtualIP(host, token string, data VirtualIPData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/system/vip", host)

	log.Printf("Creating Virtual IP: %s\n", data.Name)
	return sendRequest("POST", url, token, data)
}

func createNewServerPool(host, token string, data ServerPoolData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool", host)

	log.Printf("Creating new server pool: %s\n", data.Name)
	return sendRequest("POST", url, token, data)

}

func createNewMemberPool(host, token, poolName string, data MemberPoolData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool/pserver-list?mkey=%s", host, url.QueryEscape(poolName))

	log.Printf("Creating new member pool: %s\n", data.IP)
	return sendRequest("POST", url, token, data)
}

func sendRequest(method, url, token string, data interface{}) ([]byte, error) {
	log.Printf("sendRequest Starting\n")
	log.Printf("-------------------------------------------------\n")
	log.Printf("Method: %s\n", method)
	log.Printf("URL: %s\n", url)
	log.Printf("Token: %s\n", token)
	log.Printf("Data: %+v\n", data)
	log.Printf("-------------------------------------------------\n")

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
	req.Header.Set("Content-type", "application/json")

	// Print headers
	log.Printf("-------------------------------------------------\n")
	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			log.Printf("Header: %s: %s\n", name, value)
		}
	}
	log.Printf("JSON data: %s\n", jsonData) // Print Data
	log.Printf("-------------------------------------------------\n")

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

	vipData := VirtualIPData{
		Name:      VipName,
		Vip:       VipIp,
		Interface: Interface,
	}

	poolData := ServerPoolData{
		Name:          PoolName,
		ServerBalance: ServerBalance,
		Health:        HealthCheck,
	}

	poolMembers := make([]MemberPoolData, len(PoolMemberIPs))
	for i, ip := range PoolMemberIPs {
		poolMembers[i] = MemberPoolData{IP: ip, SSL: PoolMemberSSL, Port: PoolMemberPort}
	}

	result, err := createVirtualIP(host, token, vipData)
	if err != nil {
		log.Printf("Error creating virtual IP: %v\n", err)
		return err
	}

	result, err = createNewServerPool(host, token, poolData)
	if err != nil {
		log.Printf("Error creating server pool: %v\n", err)
		return err
	}

	for _, member := range poolMembers {
		_, err := createNewMemberPool(host, token, poolData.Name, member)
		if err != nil {
			log.Printf("Error creating member pool: %v\n", err)
			return err
		}
	}

	log.Printf("End of onboardNewApplicationPolicy\n")
	return c.JSON(http.StatusOK, string(result))
}
