package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Println("Received:", string(msg))
		conn.WriteMessage(websocket.TextMessage, []byte("Echo: "+string(msg)))
	}
}

func restHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a REST API response"))
}

func main() {
	http.HandleFunc("/ws", wsHandler)    // WebSocket endpoint
	http.HandleFunc("/api", restHandler) // REST API endpoint

	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
