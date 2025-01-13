package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Message represents the structure for the JSON response
type Message struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Get query parameter "name"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	// Create the JSON response
	response := Message{Message: fmt.Sprintf("Hello, %s!", name)}

	// Convert the struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonResponse)
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Get query parameter "name"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	// Create the JSON response
	response := Message{Message: fmt.Sprintf("Goodbye, %s!", name)}

	// Convert the struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonResponse)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Create the JSON response
	response := Message{Message: "Server is running"}

	// Convert the struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonResponse)
}

func main() {
	// Register the /hello route
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/goodbye", goodbyeHandler)
	http.HandleFunc("/", statusHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}