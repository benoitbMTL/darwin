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

// Data Types Struct

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

type VirtualServerData struct {
	Name string `json:"name,omitempty"`
}

type AssignVIPData struct {
	Interface string `json:"interface,omitempty"`
	Status    string `json:"status,omitempty"`
	VipName   string `json:"vip,omitempty"`
}

type Request struct {
	Data interface{} `json:"data"`
}

type XForwardedForData struct {
	Name                 string `json:"name,omitempty"`
	XForwardedForSupport string `json:"x-forwarded-for-support,omitempty"`
}

type ProtectionProfileData struct {
	SignatureRule     string `json:"signature-rule,omitempty"`
	XForwardedForRule string `json:"x-forwarded-for-rule,omitempty"`
}

type PolicyData struct {
	Name                    string `json:"name,omitempty"`
	DeploymentMode          string `json:"deployment-mode,omitempty"`
	Protocol                string `json:"protocol,omitempty"`
	Ssl                     string `json:"ssl,omitempty"`
	ImplicitSsl             string `json:"implicit_ssl,omitempty"`
	Vserver                 string `json:"vserver,omitempty"`
	Service                 string `json:"service,omitempty"`
	InlineProtectionProfile string `json:"web-protection-profile,omitempty"`
	ServerPool              string `json:"server-pool,omitempty"`
	TrafficLog              string `json:"tlog,omitempty"`
	HttpsService            string `json:"https-service,omitempty"`
	Certificate             string `json:"certificate,omitempty"`
}

// Virtual IP

func createNewVirtualIP(host, token string, data VirtualIPData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/system/vip", host)

	// log.Printf("Creating Virtual IP: %s\n", data.Name)
	return sendRequest("POST", url, token, data)
}

func deleteVirtualIP(host, token, vipName string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/system/vip?mkey=%s", host, url.QueryEscape(vipName))

	// log.Printf("Deleting Virtual IP: %s\n", vipName)
	return sendRequest("DELETE", url, token, nil)
}

// Server Pool

func createNewServerPool(host, token string, data ServerPoolData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool", host)

	// log.Printf("Creating new server pool: %s\n", data.Name)
	return sendRequest("POST", url, token, data)

}

func deleteServerPool(host, token, poolName string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool?mkey=%s", host, url.QueryEscape(poolName))

	// log.Printf("Deleting Server Pool: %s\n", poolName)
	return sendRequest("DELETE", url, token, nil)
}

// Member Pool

func createNewMemberPool(host, token, poolName string, data MemberPoolData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/server-pool/pserver-list?mkey=%s", host, url.QueryEscape(poolName))

	// log.Printf("Creating new member pool: %s\n", data.IP)
	return sendRequest("POST", url, token, data)
}

// Virtual Server

func createNewVirtualServer(host, token string, data VirtualServerData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/vserver", host)

	return sendRequest("POST", url, token, data)
}

func deleteVirtualServer(host, token, virtualServerName string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/vserver?mkey=%s", host, url.QueryEscape(virtualServerName))

	return sendRequest("DELETE", url, token, nil)
}

// Assign VIP to Virtual Server

func assignVIPToVirtualServer(host, token, virtualServerName string, data AssignVIPData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/vserver/vip-list?mkey=%s", host, url.QueryEscape(virtualServerName))

	// log.Printf("Assigning VIP: %s to Virtual Server: %s\n", data.VipName, virtualServerName)
	return sendRequest("POST", url, token, data)
}

// Signature Protection

func cloneSignatureProtection(host, token, originalKey, cloneKey string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/signature?mkey=%s&clone_mkey=%s", host, url.QueryEscape(originalKey), url.QueryEscape(cloneKey))

	// log.Printf("Cloning Signature Protection: %s to %s\n", originalKey, cloneKey)
	return sendRequest("POST", url, token, nil)
}

func deleteSignatureProtection(host, token, signatureName string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/signature?mkey=%s", host, url.QueryEscape(signatureName))

	// log.Printf("Deleting Signature Protection: %s\n", signatureName)
	return sendRequest("DELETE", url, token, nil)
}

// Inline Protection Profile

func cloneInlineProtection(host, token, originalKey, cloneKey string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/web-protection-profile.inline-protection?mkey=%s&clone_mkey=%s", host, url.QueryEscape(originalKey), url.QueryEscape(cloneKey))

	// log.Printf("Cloning Inline Protection: %s to %s\n", originalKey, cloneKey)
	return sendRequest("POST", url, token, nil)
}

func deleteInlineProtection(host, token, profileName string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/web-protection-profile.inline-protection?mkey=%s", host, url.QueryEscape(profileName))

	// log.Printf("Deleting Protection Profile: %s\n", profileName)
	return sendRequest("DELETE", url, token, nil)
}

// X-Forwarded-For Rule

func createNewXForwardedForRule(host, token string, data XForwardedForData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/x-forwarded-for", host)

	// log.Printf("Creating new X-Forwarded-For Rule: %s\n", data.Name)
	return sendRequest("POST", url, token, data)
}

func deleteXForwardedForRule(host, token, ruleName string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/x-forwarded-for?mkey=%s", host, url.QueryEscape(ruleName))

	// log.Printf("Deleting X-Forwarded-For Rule: %s\n", ruleName)
	return sendRequest("DELETE", url, token, nil)
}

// Protection Profile

func configureProtectionProfile(host, token, mkey string, data ProtectionProfileData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/waf/web-protection-profile.inline-protection?mkey=%s", host, url.QueryEscape(mkey))

	// log.Printf("Configuring Protection Profile: %s\n", mkey)
	return sendRequest("PUT", url, token, data)
}

// Policy

func createNewPolicy(host, token string, data PolicyData) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/policy", host)

	// log.Printf("Creating new Policy: %s\n", data.Name)
	return sendRequest("POST", url, token, data)
}

func deletePolicy(host, token, policyName string) ([]byte, error) {
	url := fmt.Sprintf("https://%s/api/v2.0/cmdb/server-policy/policy?mkey=%s", host, url.QueryEscape(policyName))

	// log.Printf("Deleting Policy: %s\n", policyName)
	return sendRequest("DELETE", url, token, nil)
}

///////////////////////////////////////////////////////////////////////////////
// Send Request                                                              //
///////////////////////////////////////////////////////////////////////////////

func sendRequest(method, url, token string, data interface{}) ([]byte, error) {
	var req *http.Request
	var err error

	reqData := Request{
		Data: data,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		// log.Printf("Error marshalling request data: %v\n", err)
		return nil, err
	}

	// Convert jsonData to string for comparison
	jsonDataStr := string(jsonData)

	if jsonDataStr != "" && jsonDataStr != `{"data":null}` {
		// Create a new request with JSON data
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	} else {
		// Create a new request without data
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		// log.Printf("Error creating HTTP request: %v\n", err)
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
		// log.Printf("Error sending HTTP request: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	// DEBUG
	// log.Printf("-------------------------------------------------\n")
	// log.Printf("sendRequest Starting\n")
	// log.Printf("URL: %s\n", url)
	// log.Printf("Method: %s\n", method)
	// Print headers
	//for name, values := range req.Header {
	// Loop over all values for the name.
	//for _, value := range values {
	// log.Printf("Header: %s: %s\n", name, value)
	//}
	//}
	// log.Printf("JSON data: %s\n", jsonData)

	time.Sleep(time.Duration(1000) * time.Millisecond)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// log.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	// log.Printf("Response received: %s\n", string(body))
	// log.Printf("-------------------------------------------------\n")

	return body, nil
}

// Checks JSON result to see if the operation was successful or not.
// It returns true if the operation was successful, and false otherwise.
func checkOperationStatus(result []byte) bool {
	var res map[string]interface{}
	json.Unmarshal(result, &res)

	// Print JSON Result
	//jsonData, err := json.Marshal(res)
	//if err != nil {
	//	log.Println(res)
	//} else {
	// log.Printf("Result JSON: %s\n", jsonData)
	//}

	// Check if the result map is empty
	if len(res) == 0 {
		// log.Printf("Operation failed: received empty result\n") // Print a message indicating that the operation failed
		return false
	}

	// Check if the "data" field is null
	if data, ok := res["data"]; ok && data == nil {
		// log.Printf("Operation failed: data field is null\n") // Print a message indicating that the operation failed
		return false
	}

	if _, ok := res["results"].(map[string]interface{})["errcode"]; ok {
		// The result contains an error code, so the operation failed
		// log.Printf("Operation failed: received an error code\n") // Print a message indicating that the operation failed
		return false
	}
	// The operation succeeded
	// log.Printf("Operation succeeded\n") // Print a message indicating that the operation succeeded
	return true
}

func calculateToken() string {
	tokenData := fmt.Sprintf(`{"username":"%s","password":"%s","vdom":"%s"}`, USERNAME_API, PASSWORD_API, VDOM_API)
	return base64.StdEncoding.EncodeToString([]byte(tokenData))
}

func onboardNewApplicationPolicy(c echo.Context) error {
	host := FWB_MGT_IP
	token := calculateToken()
	// // log.Printf("Token: %s\n", token)

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

	vsData := VirtualServerData{
		Name: VirtualServerName,
	}

	assignVIPData := AssignVIPData{
		Interface: Interface,
		Status:    VipStatus,
		VipName:   VipName,
	}

	xffData := XForwardedForData{
		Name:                 XForwardedForName,
		XForwardedForSupport: XForwardedForSupport,
	}

	protectionProfileData := ProtectionProfileData{
		SignatureRule:     CloneSignatureProtectionName,
		XForwardedForRule: XForwardedForName,
	}

	policyData := PolicyData{
		Name:                    PolicyName,
		DeploymentMode:          PolicyDeploymentMode,
		Protocol:                PolicyProtocol,
		Ssl:                     PolicySSL,
		ImplicitSsl:             PolicyImplicitSSL,
		Vserver:                 PolicyVirtualServer,
		Service:                 PolicyService,
		InlineProtectionProfile: PolicyInlineProtectionProfile,
		ServerPool:              PolicyServerPool,
		TrafficLog:              PolicyTrafficLog,
		HttpsService:            PolicyHTTPSService,
		//Certificate:          PolicyCertificate,
	}

	// Initialize a slice to store the statuses
	statuses := []map[string]string{}

	// Step 1: Create new Virtual IP
	result, err := createNewVirtualIP(host, token, vipData)
	if err != nil {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewVirtualIP",
			"status":      "failure",
			"description": "Create new Virtual IP",
			"message":     fmt.Sprintf("Error creating virtual IP: %v", err),
		})
	} else if !checkOperationStatus(result) {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewVirtualIP",
			"status":      "failure",
			"description": "Create new Virtual IP",
			"message":     "Failed to create virtual IP",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewVirtualIP",
			"status":      "success",
			"description": "Create new Virtual IP",
			"message":     "Successfully created virtual IP",
		})
	}

	// Step 2: Create new Server Pool
	result, err = createNewServerPool(host, token, poolData)
	if err != nil {
		log.Printf("Error creating server pool: %v\n", err)
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewServerPool",
			"status":      "failure",
			"description": "Create new Server Pool",
			"message":     fmt.Sprintf("Error creating Server Pool: %v", err),
		})
	} else if !checkOperationStatus(result) {
		// log.Printf("Failed to create Server Pool\n")
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewServerPool",
			"status":      "failure",
			"description": "Create new Server Pool",
			"message":     "Failed to create Server Pool",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewServerPool",
			"status":      "success",
			"description": "Create new Server Pool",
			"message":     "Successfully created Server Pool",
		})
	}

	// Step 3: Create new Member Pool
	for _, member := range poolMembers {
		result, err := createNewMemberPool(host, token, poolData.Name, member)
		if err != nil {
			statuses = append(statuses, map[string]string{
				"taskId":      "createNewMemberPool",
				"status":      "failure",
				"description": "Create new Member Pool",
				"message":     fmt.Sprintf("Error creating Member Pool: %v", err),
			})
		} else if !checkOperationStatus(result) {
			// log.Printf("Failed to create Member Pool\n")
			statuses = append(statuses, map[string]string{
				"taskId":      "createNewMemberPool",
				"status":      "failure",
				"description": "Create new Member Pool",
				"message":     "Failed to create Member Pool",
			})
		} else {
			statuses = append(statuses, map[string]string{
				"taskId":      "createNewMemberPool",
				"status":      "success",
				"description": "Create new Member Pool",
				"message":     "Successfully created Member Pool",
			})
		}
	}

	// Step 4: Create new Virtual Server
	result, err = createNewVirtualServer(host, token, vsData)
	if err != nil {
		// log.Printf("Error creating virtual server: %v\n", err)
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewVirtualServer",
			"status":      "failure",
			"description": "Create new Virtual Server",
			"message":     fmt.Sprintf("Error creating Virtual Server: %v", err),
		})
	} else if !checkOperationStatus(result) {
		// log.Printf("Failed to create Virtual Server\n")
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewVirtualServer",
			"status":      "failure",
			"description": "Create new Virtual Server",
			"message":     "Failed to create Virtual Server",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewVirtualServer",
			"status":      "success",
			"description": "Create new Virtual Server",
			"message":     "Successfully created Virtual Server",
		})
	}

	// Step 5: Assign Virtual IP to Virtual Server
	result, err = assignVIPToVirtualServer(host, token, VirtualServerName, assignVIPData)
	if err != nil {
		statuses = append(statuses, map[string]string{
			"taskId":      "assignVIPToVirtualServer",
			"status":      "failure",
			"description": "Assign Virtual IP to Virtual Server",
			"message":     fmt.Sprintf("Error assigning VIP to Virtual Server: %v", err),
		})
	} else if !checkOperationStatus(result) {
		statuses = append(statuses, map[string]string{
			"taskId":      "assignVIPToVirtualServer",
			"status":      "failure",
			"description": "Assign Virtual IP to Virtual Server",
			"message":     "Failed to assign VIP to Virtual Server",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "assignVIPToVirtualServer",
			"status":      "success",
			"description": "Assign Virtual IP to Virtual Server",
			"message":     "Successfully assigned VIP to Virtual Server",
		})
	}

	// Setp 6: Clone Signature Protection
	result, err = cloneSignatureProtection(host, token, OriginalSignatureProtectionName, CloneSignatureProtectionName)
	if err != nil {
		statuses = append(statuses, map[string]string{
			"taskId":      "cloneSignatureProtection",
			"status":      "failure",
			"description": "Clone Signature Protection",
			"message":     fmt.Sprintf("Error cloning Signature Protection: %v", err),
		})
	} else if !checkOperationStatus(result) {
		statuses = append(statuses, map[string]string{
			"taskId":      "cloneSignatureProtection",
			"status":      "failure",
			"description": "Clone Signature Protection",
			"message":     "Failed to clone Signature Protection",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "cloneSignatureProtection",
			"status":      "success",
			"description": "Clone Signature Protection",
			"message":     "Successfully cloned Signature Protection",
		})
	}

	// Setp 7: Clone Inline Protection
	result, err = cloneInlineProtection(host, token, OriginalInlineProtectionProfileName, CloneInlineProtectionProfileName)
	if err != nil {
		statuses = append(statuses, map[string]string{
			"taskId":      "cloneInlineProtection",
			"status":      "failure",
			"description": "Clone Inline Protection",
			"message":     fmt.Sprintf("Error cloning Inline Protection: %v", err),
		})
	} else if !checkOperationStatus(result) {
		statuses = append(statuses, map[string]string{
			"taskId":      "cloneInlineProtection",
			"status":      "failure",
			"description": "Clone Inline Protection",
			"message":     "Failed to clone Inline Protection",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "cloneInlineProtection",
			"status":      "success",
			"description": "Clone Inline Protection",
			"message":     "Successfully cloned Inline Protection",
		})
	}

	// Setp 8: Create new X-Forwarded-For Rule
	result, err = createNewXForwardedForRule(host, token, xffData)
	if err != nil {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewXForwardedForRule",
			"status":      "failure",
			"description": "Create new X-Forwarded-For Rule",
			"message":     fmt.Sprintf("Error creating new X-Forwarded-For Rule: %v", err),
		})
	} else if !checkOperationStatus(result) {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewXForwardedForRule",
			"status":      "failure",
			"description": "Create new X-Forwarded-For Rule",
			"message":     "Failed to create new X-Forwarded-For Rule",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewXForwardedForRule",
			"status":      "success",
			"description": "Create new X-Forwarded-For Rule",
			"message":     "Successfully created new X-Forwarded-For Rule",
		})
	}

	// Setp 9: Configure Protection Profile
	result, err = configureProtectionProfile(host, token, CloneInlineProtectionProfileName, protectionProfileData)
	if err != nil {
		statuses = append(statuses, map[string]string{
			"taskId":      "configureProtectionProfile",
			"status":      "failure",
			"description": "Configure Protection Profile",
			"message":     fmt.Sprintf("Error configuring Protection Profile: %v", err),
		})
	} else if !checkOperationStatus(result) {
		statuses = append(statuses, map[string]string{
			"taskId":      "configureProtectionProfile",
			"status":      "failure",
			"description": "Configure Protection Profile",
			"message":     "Failed to configure Protection Profile",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "configureProtectionProfile",
			"status":      "success",
			"description": "Configure Protection Profile",
			"message":     "Successfully configured Protection Profile",
		})
	}

	// Setp 10: Create new Policy
	result, err = createNewPolicy(host, token, policyData)
	if err != nil {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewPolicy",
			"status":      "failure",
			"description": "Create new Policy",
			"message":     fmt.Sprintf("Error creating new Policy: %v", err),
		})
	} else if !checkOperationStatus(result) {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewPolicy",
			"status":      "failure",
			"description": "Create new Policy",
			"message":     "Failed to create new Policy",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "createNewPolicy",
			"status":      "success",
			"description": "Create new Policy",
			"message":     "Successfully created new Policy",
		})
	}

	// log.Printf("End of onboardNewApplicationPolicy\n")
	// Return a JSON response with the statuses of all steps
	return c.JSON(http.StatusOK, statuses)
}

func deleteApplicationPolicy(c echo.Context) error {
	host := FWB_MGT_IP
	token := calculateToken()
	//// log.Printf("Token: %s\n", token)

	// Initialize a slice to store the statuses
	statuses := []map[string]string{}

	// Step 1: Delete Policy
	result, err := deletePolicy(host, token, PolicyName)
	if err != nil {
		statuses = append(statuses, map[string]string{
			"taskId":      "deletePolicy",
			"status":      "failure",
			"description": "Delete Policy",
			"message":     fmt.Sprintf("Error deleting Policy: %v", err),
		})
	} else if !checkOperationStatus(result) {
		statuses = append(statuses, map[string]string{
			"taskId":      "deletePolicy",
			"status":      "failure",
			"description": "Delete Policy",
			"message":     "Failed to delete Policy",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "deletePolicy",
			"status":      "success",
			"description": "Delete Policy",
			"message":     "Successfully deleted Policy",
		})
	}

	// Step 2: Delete Inline Protection Profile
	result, err = deleteInlineProtection(host, token, CloneInlineProtectionProfileName)
	if err != nil {
		// ...
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteInlineProtection",
			"status":      "failure",
			"description": "Delete Inline Protection Profile",
			"message":     fmt.Sprintf("Error deleting Inline Protection Profile: %v", err),
		})
	} else if !checkOperationStatus(result) {
		// ...
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteInlineProtection",
			"status":      "failure",
			"description": "Delete Inline Protection Profile",
			"message":     "Failed to delete Inline Protection Profile",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteInlineProtection",
			"status":      "success",
			"description": "Delete Inline Protection Profile",
			"message":     "Successfully deleted Inline Protection Profile",
		})
	}

	// Step 3: Delete X-Forwarded-For Rule
	result, err = deleteXForwardedForRule(host, token, XForwardedForName)
	if err != nil {
		// ...
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteXForwardedForRule",
			"status":      "failure",
			"description": "Delete X-Forwarded-For Rule",
			"message":     fmt.Sprintf("Error deleting X-Forwarded-For Rule: %v", err),
		})
	} else if !checkOperationStatus(result) {
		// ...
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteXForwardedForRule",
			"status":      "failure",
			"description": "Delete X-Forwarded-For Rule",
			"message":     "Failed to delete X-Forwarded-For Rule",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteXForwardedForRule",
			"status":      "success",
			"description": "Delete X-Forwarded-For Rule",
			"message":     "Successfully deleted X-Forwarded-For Rule",
		})
	}

	// Step 4: Delete Signature Protection
	result, err = deleteSignatureProtection(host, token, CloneSignatureProtectionName)
	if err != nil {
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteSignatureProtection",
			"status":      "failure",
			"description": "Delete Signature Protection",
			"message":     fmt.Sprintf("Error deleting Signature Protection: %v", err),
		})
	} else if !checkOperationStatus(result) {
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteSignatureProtection",
			"status":      "failure",
			"description": "Delete Signature Protection",
			"message":     "Failed to delete Signature Protection",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteSignatureProtection",
			"status":      "success",
			"description": "Delete Signature Protection",
			"message":     "Successfully deleted Signature Protection",
		})
	}

	// Step 5: DeleteVirtualServer
	result, err = deleteVirtualServer(host, token, VirtualServerName)
	if err != nil {
		// log.Printf("Error creating virtual server: %v\n", err)
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteVirtualServer",
			"status":      "failure",
			"description": "Delete Virtual Server",
			"message":     fmt.Sprintf("Error deleting Virtual Server: %v", err),
		})
	} else if !checkOperationStatus(result) {
		// log.Printf("Failed to delete Virtual Server\n")
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteVirtualServer",
			"status":      "failure",
			"description": "Delete Virtual Server",
			"message":     "Failed to delete Virtual Server",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteVirtualServer",
			"status":      "success",
			"description": "Delete Virtual Server",
			"message":     "Successfully deleted Virtual Server",
		})
	}

	// Step : Delete Server Pool
	result, err = deleteServerPool(host, token, PoolName)
	if err != nil {
		// log.Printf("Error deleting server pool: %v\n", err)
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteServerPool",
			"status":      "failure",
			"description": "Delete Server Pool",
			"message":     fmt.Sprintf("Error deleting Server Pool: %v", err),
		})
	} else if !checkOperationStatus(result) {
		// log.Printf("Failed to delete Server Pool\n")
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteServerPool",
			"status":      "failure",
			"description": "Delete Server Pool",
			"message":     "Failed to delete Server Pool",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteServerPool",
			"status":      "success",
			"description": "Delete Server Pool",
			"message":     "Successfully deleted Server Pool",
		})
	}

	// Step 7: Delete Virtual IP
	result, err = deleteVirtualIP(host, token, VipName)
	if err != nil {
		// log.Printf("Error deleting virtual IP: %v\n", err)
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteVirtualIP",
			"status":      "failure",
			"description": "Delete Virtual IP",
			"message":     fmt.Sprintf("Error deleting virtual IP: %v", err),
		})
	} else if !checkOperationStatus(result) {
		// log.Printf("Failed to delete Virtual IP\n")
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteVirtualIP",
			"status":      "failure",
			"description": "Delete Virtual IP",
			"message":     "Failed to delete Virtual IP",
		})
	} else {
		statuses = append(statuses, map[string]string{
			"taskId":      "deleteVirtualIP",
			"status":      "success",
			"description": "Delete Virtual IP",
			"message":     "Successfully deleted virtual IP",
		})
	}

	// log.Printf("End of deleteApplicationPolicy\n")
	// Return a JSON response with the statuses of all steps
	return c.JSON(http.StatusOK, statuses)
}
