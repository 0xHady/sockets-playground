package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Close the connection when the function returns
	defer ws.Close()

	log.Println("Client connected")

	const (
		// Time allowed to read the next pong message from the peer
		pongWait = 2 * time.Second

		// Maximum message size allowed from peer
		maxMessageSize = 512
	)

	// Listen for incoming messages
	for {
		// Read message from browser
		ws.SetReadLimit(maxMessageSize)
		// ws.SetReadDeadline(time.Now().Add(pongWait))

		ws.SetPongHandler(func(string) error {
			ws.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Received: %s\n", p)

		// Echo the message back to the client
		if err := ws.WriteMessage(messageType, []byte(`fuck off`)); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Configure WebSocket route
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("New WebSocket connection attempt")
		handleConnections(w, r)
	})

	// Start the server on port 8080
	addr := "localhost:8080"
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server starting on http://%s\n", addr)
	// log.Printf("Ping interval: %v, Pong timeout: %v\n", pingPeriod, pongWait)
	log.Fatal(server.ListenAndServe())
}
