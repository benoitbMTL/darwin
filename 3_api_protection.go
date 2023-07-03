package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	fmt.Println("Status:", status) // Debug status

	apiURL := fmt.Sprintf("%s/%s", PETSTORE_URL, status)
	fmt.Println("API URL:", apiURL) // Debug API URL

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

	body, _ := ioutil.ReadAll(resp.Body)

	// Debug Body Response
	fmt.Println("Response Body:", string(body))

	var js json.RawMessage
	if json.Unmarshal(body, &js) != nil {
		// If it's not JSON, return the raw body
		return c.JSON(http.StatusOK, map[string]string{"body": string(body)})
	}

	var pets PetstorePet
	err = json.Unmarshal(body, &pets)
	if err != nil {
		fmt.Println("Unmarshal Error:", err) // Debug unmarshal error
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	fmt.Println("Unmarshalled Pets:", pets) // Debug unmarshalled result

	return c.JSON(http.StatusOK, pets)
}


func handlePetstoreAPIRequestPost(c echo.Context) error {
	return nil
}

func handlePetstoreAPIRequestPut(c echo.Context) error {
	return nil
}

func handlePetstoreAPIRequestDelete(c echo.Context) error {
	return nil
}
