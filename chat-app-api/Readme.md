# Real-Time Chat Application

A Go-based backend for a real-time chat application using WebSockets. This application supports multiple clients in a chat room and broadcasts messages in real time.

---

## Features

### Real-Time Communication:

Persistent, two-way communication using WebSockets.
Broadcasts messages to all connected clients in a chat room.

### Concurrency:

Handles multiple clients simultaneously using Go’s goroutines and channels.

### Thread-Safe State Management:

Safely manages client connections using sync.Mutex.
Technologies Used
Programming Language: Go (Golang)
WebSocket Library: gorilla/websocket

---

### Architecture Overview

#### ChatRoom:

Manages all active clients and broadcasts messages.
Uses:
A map to store active clients.
A channel to handle messages.
Client:

Represents a single WebSocket connection.
Handles:
Reading messages from the WebSocket.
Writing messages to the WebSocket.
WebSocket Handler:

Upgrades HTTP requests to WebSocket connections.
Registers new clients and starts their Read and Write goroutines.
Project Structure
bash
Copy code
chat-app/
├── main.go # Entry point of the application
├── chatroom.go # ChatRoom struct and message broadcasting logic
├── client.go # Client struct and WebSocket read/write logic
├── handlers.go # WebSocket handler to manage client connections
└── go.mod # Module file for dependency management
Installation and Setup

1. Clone the Repository
   bash
   Copy code
   git clone <repository-url>
   cd chat-app
2. Install Dependencies
   Ensure you have Go installed. Then, install the required WebSocket library:

bash
Copy code
go mod tidy 3. Run the Server
Start the WebSocket server:

bash
Copy code
go run main.go
The server will start on http://localhost:8080.

API Endpoints
WebSocket Endpoint
Route: /ws
Method: WebSocket connection
Description: Upgrades HTTP connections to WebSocket connections and joins the chat room.
Testing the Application

1. Start the Server
   Run the following command:

bash
Copy code
go run main.go 2. Connect Clients
Use a WebSocket testing tool like WebSocket King.
Connect multiple clients to:
bash
Copy code
ws://localhost:8080/ws 3. Send Messages
From one client, send a message:
text
Copy code
Hello from Client 1
Verify that all connected clients receive the message.
Key Concepts
WebSockets:

Persistent, full-duplex communication for real-time applications.
Allows bidirectional communication between clients and the server.
Concurrency:

Used Go’s goroutines to handle Read and Write operations for multiple clients simultaneously.
Synchronization:

Managed shared state (client list) with sync.Mutex for thread-safe updates.
Channels:

Used Go channels to broadcast messages to all connected clients.
Future Enhancements
Multiple Chat Rooms:

Support separate chat rooms (e.g., /ws/room1, /ws/room2).
User Authentication:

Use JWT tokens to authenticate users before joining.
Message Persistence:

Store chat messages in a database for retrieval.
Message Formatting:

Add timestamps or sender details to messages.
Private Messaging:

Support direct messages between specific users.
