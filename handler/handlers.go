package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

func ModelsHandler(w http.ResponseWriter, r *http.Request) {
	models := []string{"gemini-2.5-flash-preview", "gemini-2.5-pro-preview", "gemini-2.0-flash"}
	
	type ModelData struct {
		ID      string `json:"id"`
		Object  string `json:"object"`
		Created int64  `json:"created"`
		OwnedBy string `json:"owned_by"`
	}

	var data []ModelData
	now := time.Now().Unix()
	for _, m := range models {
		data = append(data, ModelData{
			ID:      m,
			Object:  "model",
			Created: now,
			OwnedBy: "google",
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"object": "list",
		"data":   data,
	})
}