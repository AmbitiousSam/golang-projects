package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader configuration
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// chatHandler upgrades the connection to WebSocket and manages clients
func chatHandler(room *ChatRoom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
			return
		}

		client := &Client{
			socket: conn,
			room:   room,
			send:   make(chan string),
		}

		room.Join(client)

		// Start client Read and Write routines
		go client.Read()
		go client.Write()
	}
}
