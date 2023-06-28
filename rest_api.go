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

func deleteVirtualIP(host, token, vipName string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/system/vip?mkey=%s", host, url.QueryEscape(vipName))

	log.Printf("Deleting Virtual IP: %s\n", vipName)
	return sendRequest("DELETE", url, token, nil)
}

func createNewServerPool(host, token string, data ServerPoolData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool", host)

	log.Printf("Creating new server pool: %s\n", data.Name)
	return sendRequest("POST", url, token, data)

}

func deleteServerPool(host, token, poolName string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool?mkey=%s", host, url.QueryEscape(poolName))

	log.Printf("Deleting Server Pool: %s\n", poolName)
	return sendRequest("DELETE", url, token, nil)
}

func createNewMemberPool(host, token, poolName string, data MemberPoolData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool/pserver-list?mkey=%s", host, url.QueryEscape(poolName))

	log.Printf("Creating new member pool: %s\n", data.IP)
	return sendRequest("POST", url, token, data)
}

func sendRequest(method, url, token string, data interface{}) ([]byte, error) {
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

	log.Printf("-------------------------------------------------\n")
	log.Printf("sendRequest Starting\n")
	log.Printf("URL: %s\n", url)
	log.Printf("Method: %s\n", method)
	// Print headers
	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			log.Printf("Header: %s: %s\n", name, value)
		}
	}
	log.Printf("JSON data: %s\n", jsonData)

	time.Sleep(time.Duration(500) * time.Millisecond)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	log.Printf("Response received: %s\n", string(body))
	log.Printf("-------------------------------------------------\n")

	return body, nil
}

// checkOperationStatus checks if the operation was successful or not.
// It returns true if the operation was successful, and false otherwise.
func checkOperationStatus(result []byte) bool {
	var res map[string]interface{}
	json.Unmarshal(result, &res)
	if _, ok := res["results"].(map[string]interface{})["errcode"]; ok {
		// The result contains an error code, so the operation failed
		return false
	}
	// The operation succeeded
	return true
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
		// Send a message to the frontend to update the task status to 'failure'
		return c.JSON(http.StatusOK, map[string]string{
			"taskId":  "createVirtualIP",
			"status":  "failure",
			"message": fmt.Sprintf("Error creating virtual IP: %v", err),
		})
	} else if !checkOperationStatus(result) {
		log.Printf("Failed to create virtual IP\n")
		// Send a message to the frontend to update the task status to 'failure'
		return c.JSON(http.StatusOK, map[string]string{
			"taskId":  "createVirtualIP",
			"status":  "failure",
			"message": "Failed to create virtual IP",
		})
	}

	result, err = createNewServerPool(host, token, poolData)
	if err != nil {
		log.Printf("Error creating server pool: %v\n", err)
		// Send a message to the frontend to update the task status to 'failure'
		return c.JSON(http.StatusOK, map[string]string{
			"taskId":  "createNewServerPool",
			"status":  "failure",
			"message": fmt.Sprintf("Error creating Server Pool: %v", err),
		})
	} else if !checkOperationStatus(result) {
		log.Printf("Failed to create Server Pool\n")
		// Send a message to the frontend to update the task status to 'failure'
		return c.JSON(http.StatusOK, map[string]string{
			"taskId":  "createNewServerPool",
			"status":  "failure",
			"message": "Failed to create Server Pool",
		})
	}

	for _, member := range poolMembers {
		result, err := createNewMemberPool(host, token, poolData.Name, member)
		if err != nil {
			// If there's an error, return a JSON response with status "failure" and an error message
			return c.JSON(http.StatusOK, map[string]string{
				"taskId":  "createNewMemberPool",
				"status":  "failure",
				"message": fmt.Sprintf("Error creating Member Pool: %v", err),
			})
		} else if !checkOperationStatus(result) {
			log.Printf("Failed to create Member Pool\n")
			// Send a message to the frontend to update the task status to 'failure'
			return c.JSON(http.StatusOK, map[string]string{
				"taskId":  "createNewMemberPool",
				"status":  "failure",
				"message": "Failed to create Member Pool",
			})
		}
	}

	log.Printf("End of onboardNewApplicationPolicy\n")
	// If everything is successful, return a JSON response with status "success" and a success message
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "End of onboardNewApplicationPolicy",
	})
}

func deleteApplicationPolicy(c echo.Context) error {
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

	// Delete server pool
	_, err := deleteServerPool(host, token, poolData.Name)
	if err != nil {
		log.Printf("Error deleting server pool: %v\n", err)
		return err
	}

	// Delete virtual IP
	_, err = deleteVirtualIP(host, token, vipData.Name)
	if err != nil {
		log.Printf("Error deleting virtual IP: %v\n", err)
		return err
	}

	log.Printf("End of deleteApplicationPolicy\n")
	return c.NoContent(http.StatusOK)
}
