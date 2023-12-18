package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type PetstorePet []struct {
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
		fmt.Println("HTTP Request Error:", err) // Debug HTTP request error
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Debug Body Response
	// fmt.Println("Response Body:", string(body))

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var pets PetstorePet
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

	// Read the request body for debugging
	PetData, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Print debug
	fmt.Println("Request URL:", apiURL)
	fmt.Println("Request Body:", string(PetData))

	// Since the original body is now consumed, create a new io.Reader from the read bytes
	reqBody := bytes.NewReader(PetData)

	req, err := http.NewRequest("POST", apiURL, reqBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

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
		fmt.Println("HTTP Request Error:", err) // Debug HTTP request error
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Debug Body Response
	fmt.Println("Response Body:", string(body))

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var pets PetstorePet
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

func handlePetstoreAPIRequestPut(c echo.Context) error {
	return nil
}

func handlePetstoreAPIRequestDelete(c echo.Context) error {
	return nil
}
