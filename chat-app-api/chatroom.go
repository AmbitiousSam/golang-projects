package main

import (
	"fmt"
	"sync"
)

type ChatRoom struct {
	clients   map[*Client]bool
	broadcast chan string
	lock      sync.Mutex
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		clients:   make(map[*Client]bool),
		broadcast: make(chan string),
	}
}

func (room *ChatRoom) Join(client *Client) {
	room.lock.Lock()
	defer room.lock.Unlock()
	room.clients[client] = true
	fmt.Println("New client joined!")
}

func (room *ChatRoom) Leave(client *Client) {
	room.lock.Lock()
	defer room.lock.Unlock()
	delete(room.clients, client)
	fmt.Println("New client joined!")
}

func (room *ChatRoom) Start() {
	for message := range room.broadcast {
		room.lock.Lock()
		for client := range room.clients {
			client.send <- message
		}
		room.lock.Unlock()
	}
}
