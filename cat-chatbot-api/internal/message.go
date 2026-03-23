package internal

import (
	"encoding/json"
    "net/http"
    "strings"
)

type ChatRequest struct {
   	Message string `json:"message"`
    UserId  string `json:"user_id"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

 	var req ChatRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

	if strings.TrimSpace(req.Message) == "" {
		http.Error(w, "Message is required", http.StatusBadRequest)
		return
	}

	resp := GetBotReply(req.Message)

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)

}