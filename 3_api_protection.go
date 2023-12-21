package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

///////////////////////////////////////////////////////////////////////////////////
// STRUCTURE                                                                     //
///////////////////////////////////////////////////////////////////////////////////

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PetstorePet struct {
	ID        int      `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags"`
	Status    string   `json:"status"`
}

type PetstorePetArray []PetstorePet

///////////////////////////////////////////////////////////////////////////////////
// GET                                                                           //
///////////////////////////////////////////////////////////////////////////////////

func handlePetstoreAPIRequestGet(c echo.Context) error {
	status := c.FormValue("status")
	// fmt.Println("Status:", status) // Debug status

	apiURL := fmt.Sprintf("%s/%s", PETSTORE_URL, status)
	// fmt.Println("API URL:", apiURL) // Debug API URL

	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Create a custom http.Client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		// fmt.Println("HTTP Request Error:", err) // Debug HTTP request error
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Debug Body Response
	// fmt.Println("Response Body:", string(body))

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var pets PetstorePetArray
		err = json.Unmarshal(body, &pets)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		// Return a JSON
		return c.JSON(http.StatusOK, pets)
	} else if strings.Contains(contentType, "text/plain") {
		// Return a TEXT
		return c.String(http.StatusOK, string(body))
	} else if strings.Contains(contentType, "text/html") {
		// Return HTML
		return c.HTML(http.StatusOK, string(body))
	} else {
		// Return an Error
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "unsupported content type"})
	}
}

///////////////////////////////////////////////////////////////////////////////////
// POST                                                                          //
///////////////////////////////////////////////////////////////////////////////////

func handlePetstoreAPIRequestPost(c echo.Context) error {
	apiURL := PETSTORE_URL

	// Read the request body
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		// Handle error if reading the request body fails
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer c.Request().Body.Close()

    // Remove the double quotes if present
    trimmedBody := strings.Trim(string(body), "\"")

    fmt.Println("Trimmed Data:", trimmedBody)

    // Check if the trimmed body is Base64 encoded
    decodedBody, decodeErr := base64.StdEncoding.DecodeString(trimmedBody)
    if decodeErr == nil {
        // It's Base64, so use the decoded body
        body = decodedBody
        fmt.Println("Decoded Base64 Data:", string(body))
    } else {
        fmt.Println("Received Regular JSON Data:", trimmedBody)
    }

    // Attempt to unmarshal the data
    var data PetstorePet
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println("Error Unmarshalling Data:", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    fmt.Printf("Unmarshalled Data: %+v\n", data)


	// Create a new POST request using the original body
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(originalBody))
	if err != nil {
		// Handle error if new request creation fails
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Set headers for the request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Print the request URL for debugging
	// fmt.Println("Request URL:", apiURL)

	// Create a custom HTTP client with a specific transport configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Skip TLS certificate verification
			},
		},
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		// Handle error if the request fails
		// fmt.Println("HTTP Request Error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	// Print the HTTP response status code
	// fmt.Println("HTTP Response Code:", resp.StatusCode)

	// Read the response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		// Handle error if reading the response body fails
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Print the response body for debugging
	// fmt.Println("Response Body:", string(body))

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var pets PetstorePet
		err = json.Unmarshal(body, &pets)
		if err != nil {
			// Handle error if JSON unmarshalling fails
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		// Return a JSON response
		return c.JSON(http.StatusOK, pets)
	} else if strings.Contains(contentType, "text/plain") {
		// Return a plain text response
		return c.String(http.StatusOK, string(body))
	} else if strings.Contains(contentType, "text/html") {
		// Return an HTML response
		return c.HTML(http.StatusOK, string(body))
	} else {
		// Handle unsupported content types
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "unsupported content type"})
	}
}

///////////////////////////////////////////////////////////////////////////////////
// PUT                                                                           //
///////////////////////////////////////////////////////////////////////////////////

func handlePetstoreAPIRequestPut(c echo.Context) error {
	apiURL := PETSTORE_URL

	// Create a new PUT request using the body from the incoming request
	req, err := http.NewRequest("PUT", apiURL, c.Request().Body)
	if err != nil {
		// Handle error if new request creation fails
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Set headers for the request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Print the request URL for debugging
	// fmt.Println("Request URL:", apiURL)

	// Create a custom HTTP client with a specific transport configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Skip TLS certificate verification
			},
		},
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		// Handle error if the request fails
		// fmt.Println("HTTP Request Error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	// Print the HTTP response status code
	// fmt.Println("HTTP Response Code:", resp.StatusCode)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle error if reading the response body fails
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Print the response body for debugging
	// fmt.Println("Response Body:", string(body))

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var pets PetstorePet
		err = json.Unmarshal(body, &pets)
		if err != nil {
			// Handle error if JSON unmarshalling fails
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		// Return a JSON response
		return c.JSON(http.StatusOK, pets)
	} else if strings.Contains(contentType, "text/plain") {
		// Return a plain text response
		return c.String(http.StatusOK, string(body))
	} else if strings.Contains(contentType, "text/html") {
		// Return an HTML response
		return c.HTML(http.StatusOK, string(body))
	} else {
		// Handle unsupported content types
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "unsupported content type"})
	}
}

///////////////////////////////////////////////////////////////////////////////////
// DELETE                                                                        //
///////////////////////////////////////////////////////////////////////////////////

func handlePetstoreAPIRequestDelete(c echo.Context) error {

	petID := c.FormValue("petId")

	apiURL := fmt.Sprintf("%s/%s", PETSTORE_URL, petID)
	// fmt.Println("PETSTORE_URL:", PETSTORE_URL) // Debug PETSTORE_URL
	// fmt.Println("API URL:", apiURL)            // Debug API URL

	req, _ := http.NewRequest("DELETE", apiURL, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Create a custom http.Client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		// fmt.Println("HTTP Request Error:", err) // Debug HTTP request error
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error reading response body: " + err.Error()})
	}

	// Debug Body Response
	// fmt.Println("Response Body:", string(body))

	contentType := resp.Header.Get("Content-Type")
	// fmt.Println("contentType:", contentType)

	if strings.Contains(contentType, "application/json") {
		// Attempt to unmarshal as JSON
		var jsonResponse map[string]interface{}
		err := json.Unmarshal(body, &jsonResponse)
		if err != nil {
			// If unmarshalling fails, treat as plain text
			return c.String(http.StatusOK, string(body))
		}

		// Check for error code in JSON response
		if code, exists := jsonResponse["code"]; exists && code.(float64) >= 400 {
			return c.JSON(http.StatusBadRequest, jsonResponse)
		}
		return c.JSON(http.StatusOK, jsonResponse)
	} else if strings.Contains(contentType, "text/plain") || strings.Contains(contentType, "text/html") {
		return c.String(http.StatusOK, string(body))
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Unsupported content type: " + contentType})
	}
}

///////////////////////////////////////////////////////////////////////////////////
// GENERATE RANDOM API TRAFFIC                                                   //
///////////////////////////////////////////////////////////////////////////////////

// randomPublicIP generates a random public IPv4 address.
func randomPublicIP() string {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for {
		ip := net.IPv4(byte(random.Intn(256)), byte(random.Intn(256)), byte(random.Intn(256)), byte(random.Intn(256)))
		if isPublicIPv4(ip) {
			return ip.String()
		}
	}
}

// isPublicIPv4 checks if an IP address is a public IPv4 address.
func isPublicIPv4(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	for _, network := range privateIPv4Networks() {
		if network.Contains(ip) {
			return false
		}
	}
	return true
}

// privateIPv4Networks returns a slice of private (RFC1918) IPv4 networks.
func privateIPv4Networks() []*net.IPNet {
	var networks []*net.IPNet
	for _, cidr := range []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"} {
		_, network, _ := net.ParseCIDR(cidr)
		networks = append(networks, network)
	}
	return networks
}

func generateRandomValue(values []string) string {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return values[random.Intn(len(values))]
}

func sendPostRequest(petStoreURL string, userAgent string, pet PetstorePet, xForwardedFor string) error {
	jsonData, err := json.Marshal(pet)
	if err != nil {
		log.Printf("Error marshalling pet data: %v\n", err)
		return err
	}

	req, err := http.NewRequest("POST", petStoreURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Forwarded-For", xForwardedFor)

	// Print equivalent curl command
	// fmt.Printf("curl -X PUT %s -H \"Accept: application/json\" -H \"Content-Type: application/json\" -H \"User-Agent: %s\" -H \"X-Forwarded-For: %s\" -d '%s'\n", petStoreURL, userAgent, xForwardedFor, jsonData)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	// Read and log the response body (optional)
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// log.Printf("Response Body: %s\n", body)

	// Print various response fields
	// fmt.Printf("Status: %s\n", resp.Status)
	// fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	// fmt.Printf("Header: %v\n", resp.Header)
	// fmt.Printf("Body: %s\n", body)
	// fmt.Printf("ContentLength: %d\n", resp.ContentLength)
	// fmt.Printf("TransferEncoding: %v\n", resp.TransferEncoding)
	// fmt.Printf("Close: %v\n", resp.Close)
	// fmt.Printf("Uncompressed: %v\n", resp.Uncompressed)

	// Handle the response as needed
	return nil
}

func sendPutRequest(petStoreURL string, userAgent string, pet PetstorePet, xForwardedFor string) error {
	jsonData, err := json.Marshal(pet)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", petStoreURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Forwarded-For", xForwardedFor)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read and log the response body (optional)
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// log.Printf("Response Body: %s\n", body)

	// Handle the response as needed
	return nil
}

func sendGetRequest(petStoreURL, randomStatus, userAgent, xForwardedFor string) error {
	// Construct the URL with query parameters
	fullURL := fmt.Sprintf("%s/findByStatus?status=%s", petStoreURL, url.QueryEscape(randomStatus))

	// Create the request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Forwarded-For", xForwardedFor)

	// Send the request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read and log the response body (optional)
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// log.Printf("Response Body: %s\n", body)

	return nil
}

func sendDeleteRequest(petStoreURL string, randomID int, userAgent string, xForwardedFor string) error {
	// Construct the URL
	fullURL := fmt.Sprintf("%s/%d", petStoreURL, randomID)

	// Create the request
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Forwarded-For", xForwardedFor)

	// Send the request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read and log the response body (optional)
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// log.Printf("Response Body: %s\n", body)

	return nil
}

func handleAPITrafficGenerator(c echo.Context) error {
	requestCount := 1800
	petNames := []string{"FortiPuma", "FortiFish", "FortiSpider", "FortiTiger", "FortiLion", "FortiShark", "FortiSnake", "FortiMonkey", "FortiFox", "FortiRam", "FortiEagle", "FortiBee", "FortiCat", "FortiDog", "FortiAnt", "FortiWasp", "FortiPanter", "FortiGator", "FortiOwl", "FortiWildcats"}
	petTypes := []string{"Puma", "Fish", "Spider", "Tiger", "Lion", "Shark", "Snake", "Monkey", "Fox", "Ram", "Eagle", "Bee", "Cat", "Dog", "Ant", "Wasp", "Panter", "Gator", "Owl", "Wildcats"}
	statuses := []string{"available", "pending", "sold"}

	for i := 0; i < requestCount; i++ {
		randomName := generateRandomValue(petNames)
		randomPet := generateRandomValue(petTypes)
		randomStatus := generateRandomValue(statuses)
		randomStatusNew := generateRandomValue(statuses)
		randomIP := randomPublicIP()
		randomID := rand.Intn(1001)
		userAgent := USER_AGENT
		petStoreURL := PETSTORE_URL

		petNew := PetstorePet{
			ID: randomID,
			Category: Category{
				ID:   randomID,
				Name: randomPet,
			},
			Name:      randomName,
			PhotoUrls: []string{randomPet + ".png"},
			Tags: []Tag{
				{
					ID:   randomID,
					Name: randomName,
				},
			},
			Status: randomStatus,
		}

		petModified := PetstorePet{
			ID: randomID,
			Category: Category{
				ID:   randomID,
				Name: randomPet,
			},
			Name:      randomName,
			PhotoUrls: []string{randomPet + ".png"},
			Tags: []Tag{
				{
					ID:   randomID,
					Name: randomName,
				},
			},
			Status: randomStatusNew,
		}

		// Send POST request
		err := sendPostRequest(petStoreURL, userAgent, petNew, randomIP)
		if err != nil {
			log.Fatalf("Error sending POST request: %v", err)
		}

		// Send PUT request
		err = sendPutRequest(petStoreURL, userAgent, petModified, randomIP)
		if err != nil {
			log.Fatalf("Error sending PUT request: %v", err)
		}

		// Send GET request
		err = sendGetRequest(petStoreURL, randomStatus, userAgent, randomIP)
		if err != nil {
			log.Fatalf("Error sending PUT request: %v", err)
		}

		// Send DELETE request
		err = sendDeleteRequest(petStoreURL, randomID, userAgent, randomIP)
		if err != nil {
			log.Fatalf("Error sending PUT request: %v", err)
		}
	}

	// Return the completion message
	message := fmt.Sprintf("API traffic generation is complete. We have sent %d random requests of types POST, PUT, GET, and DELETE.", requestCount)
	return c.String(http.StatusOK, message)
}
