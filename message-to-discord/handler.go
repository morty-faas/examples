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

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This function accept only POST requests", http.StatusMethodNotAllowed)
		return
	}

	//TODO: get message from request
	data, _ := json.Marshal(map[string]interface{}{
		"user":    "Morty",
		"content": "Hello from a Morty function !",
	})

	req, _ := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	if _, err := http.DefaultClient.Do(req); err != nil {
		http.Error(w, fmt.Sprintf("failed to send data to webhook: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
