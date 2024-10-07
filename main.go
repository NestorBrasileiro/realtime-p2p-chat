package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer ws.Close()

	for {
		_, message, err := ws.ReadMessage()

		if err != nil {
			fmt.Println("Error reading message: ", err)
			break
		}
		fmt.Println("Received message: ", string(message))

		if err := ws.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}

	}

}

func main() {

	http.HandleFunc("/f_", handleConnections)
	fmt.Println("Server listening on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}

}
