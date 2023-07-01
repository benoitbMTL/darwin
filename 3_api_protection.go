package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

type PetstoreRequest struct {
	Status string `json:"status"`
}

func handlePetstoreRequest(w http.ResponseWriter, r *http.Request) {
	var req PetstoreRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := exec.Command("curl", "-s", "-k", "-X", "GET", "${PETSTORE_URL}/pet/findByStatus?status="+req.Status, "-H", "Accept: application/json", "-H", "Content-Type: application/json")
	cmdOutput, err := cmd.Output()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cmdOutput)
}
