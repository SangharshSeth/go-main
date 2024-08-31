package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Store WebSocket connections in a slice
var connections = make([]*websocket.Conn, 0)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSockets(w http.ResponseWriter, r *http.Request) {
	log.Print("Came to Websocket connection function")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	// Add the new connection to the list
	connections = append(connections, conn)
	fmt.Println("New WebSocket connection")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		log.Printf("Message received %s", message)
	}
}

// Broadcast message to all WebSocket clients
func BroadcastMessage(message []byte) {
	for _, conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}
