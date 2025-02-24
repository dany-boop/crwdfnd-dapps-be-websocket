package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

var (
	clients   = make(map[*Client]bool) // Active WebSocket connections
	broadcast = make(chan []byte)      // Channel for broadcasting messages
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	mu sync.Mutex
)

// Handles incoming WebSocket connections
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}

	client := &Client{Conn: conn, Send: make(chan []byte)}
	mu.Lock()
	clients[client] = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, client)
		mu.Unlock()
		conn.Close()
	}()

	go handleMessages(client)
}

// Handles incoming messages and broadcasts them
func handleMessages(client *Client) {
	defer client.Conn.Close()
	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(clients, client)
			mu.Unlock()
			break
		}
		broadcast <- msg
	}
}

// Sends messages to all connected clients
func StartBroadcaster() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			select {
			case client.Send <- msg:
				go func(c *Client) {
					c.Conn.WriteMessage(websocket.TextMessage, <-c.Send)
				}(client)
			default:
				close(client.Send)
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
