package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// GetUsersHandler is...
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"id": "12345",
		"ts": time.Now().Format(time.RFC3339),
	}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write(b)
}
