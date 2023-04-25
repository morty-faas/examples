package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	webhookUrl = "https://discord.com/api/webhooks/1089563947423764550/LuUfqOQMwscSZlS9gOboRraau0J-rJw9YCMKJ03RA2P5q1GKOGQUUB8Ezqlr1G4T3EIz"
)

// Data Transfer Object for the request
type PublishMessageRequest struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// If the incoming request is not POST,
	// we return a HTTP 405 Method Not Allowed
	if r.Method != http.MethodPost {
		http.Error(w, "This function accept only POST requests.", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body, and if an error occurs,
	// return a HTTP 400 Bad Request error
	dto := &PublishMessageRequest{}
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		http.Error(w, fmt.Sprintf("Unable to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	// Create a valid JSON request body for the webhook.
	// https://discord.com/developers/docs/resources/webhook#execute-webhook
	data, _ := json.Marshal(map[string]interface{}{
		"user":    "Morty",
		"content": dto.Message,
	})

	// Create the request and execute it
	req, _ := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	if _, err := http.DefaultClient.Do(req); err != nil {
		http.Error(w, fmt.Sprintf("failed to send data to webhook: %v", err), http.StatusInternalServerError)
		return
	}

	// Send the HTTP 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
