package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

type PetstoreRequest struct {
	Status string `json:"status"`
}

func handlePetstoreAPIRequest(c echo.Context) error {
	req := new(PetstoreRequest)

	if err := c.Bind(req); err != nil {
		log.Println("Error binding request:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	log.Println("Received status:", req.Status)

	curlCommand := "curl -s -k -X GET " + PETSTORE_URL + "/pet/findByStatus?status=" + req.Status + " -H 'Accept: application/json' -H 'Content-Type: application/json' | jq"
	log.Println("Petstore URL:", PETSTORE_URL)
	log.Println("CURL Command:", curlCommand)

	cmd := exec.Command("sh", "-c", curlCommand)
	cmdOutput, err := cmd.Output()

	if err != nil {
		log.Println("Error executing curl command:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Println("CURL Result:", string(cmdOutput))

	return c.JSONBlob(http.StatusOK, cmdOutput)
}
