package main

import (
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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	cmd := exec.Command("curl", "-s", "-k", "-X", "GET", "${PETSTORE_URL}/pet/findByStatus?status="+req.Status, "-H", "Accept: application/json", "-H", "Content-Type: application/json")
	cmdOutput, err := cmd.Output()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSONBlob(http.StatusOK, cmdOutput)
}
