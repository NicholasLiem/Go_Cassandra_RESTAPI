package handlers

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Message string `json:"message"`
}

func StatusHandler(rw http.ResponseWriter, r *http.Request) {
	response := StatusResponse{
		Message: "Hello, World!",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(jsonData)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}
}
