package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	socket *websocket.Conn
	room   *ChatRoom
	send   chan string
}

func (c *Client) Read() {
	defer func() {
		c.room.Leave(c)
		c.socket.Close()
	}()
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			break
		}
		c.room.broadcast <- string(message)
	}
}

func (c *Client) Write() {
	defer c.socket.Close()
	for message := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			break
		}
	}
}
