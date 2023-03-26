package function

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DebugResponse struct {
	Method  string                 `json:"method"`
	Path    string                 `json:"path"`
	Body    map[string]interface{} `json:"body,omitempty"`
	Headers http.Header            `json:"headers,omitempty"`
	Query   string                 `json:"query,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	data := map[string]interface{}{}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request body: %v", err), http.StatusInternalServerError)
		return
	}

	if len(body) != 0 {
		if err := json.Unmarshal(body, &data); err != nil {
			http.Error(w, fmt.Sprintf("failed to decode request body: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&DebugResponse{
		Method:  r.Method,
		Path:    r.URL.Path,
		Headers: r.Header,
		Query:   r.URL.RawQuery,
		Body:    data,
	})
}
