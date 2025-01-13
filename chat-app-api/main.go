package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new chat room
	room := NewChatRoom()

	// Start broadcasting messages
	go room.Start()

	// Register WebSocket handler
	http.HandleFunc("/ws", chatHandler(room))

	// Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
