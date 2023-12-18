package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type PetstorePet struct {
	ID        int      `json:"id,omitempty"`
	Category  Category `json:"category,omitempty"`
	Name      string   `json:"name,omitempty"`
	PhotoUrls []string `json:"photoUrls,omitempty"`
	Tags      []Tags   `json:"tags,omitempty"`
	Status    string   `json:"status,omitempty"`
}

type Category struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Tags struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PetstorePetArray []PetstorePet

func handlePetstoreAPIRequestGet(c echo.Context) error {
	status := c.FormValue("status")
	//fmt.Println("Status:", status) // Debug status

	apiURL := fmt.Sprintf("%s/%s", PETSTORE_URL, status)
	//fmt.Println("API URL:", apiURL) // Debug API URL

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

func handlePetstoreAPIRequestPost(c echo.Context) error {
    apiURL := PETSTORE_URL

    // Create a new POST request using the body from the incoming request
    req, err := http.NewRequest("POST", apiURL, c.Request().Body)
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

func handlePetstoreAPIRequestDelete(c echo.Context) error {
	petID := c.FormValue("pet-id") // Receive pet ID as a string
	fmt.Println("Pet ID:", petID) // Debug pet ID

	apiURL := fmt.Sprintf("%s/%s", PETSTORE_URL, petID)
	fmt.Println("PETSTORE_URL:", PETSTORE_URL) // Debug PETSTORE_URL
    fmt.Println("API URL:", apiURL) // Debug API URL

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

	body, _ := io.ReadAll(resp.Body)

	// Debug Body Response
	// fmt.Println("Response Body:", string(body))

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var pet PetstorePet
		err = json.Unmarshal(body, &pet)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		// Return a JSON
		return c.JSON(http.StatusOK, pet)
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

