package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request struct {
		LongURL string `json:"long_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil || request.LongURL == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	shortID := uuid.New().String()[:8]
	shortURL := "http://localhost:8080/" + shortID

	mapping := URLMapping{shortURL: shortID, LongURL: request.LongURL}
	db.Create(&mapping)
	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortID := r.URL.Path[1:]
	var mapping URLMapping
	result := db.First(&mapping, "short_url = ?", shortID)
	if result.Error != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
	}
	http.Redirect(w, r, mapping.LongURL, http.StatusFound)
}
