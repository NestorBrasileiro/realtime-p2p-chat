package services

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Mutex      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (hub *Hub) Run() {

	for {
		select {
		case client := <-hub.Register:
			hub.Mutex.Lock()
			hub.Clients[client] = true
			hub.Mutex.Unlock()
		case client := <-hub.Unregister:
			hub.Mutex.Lock()
			if _, ok := hub.Clients[client]; ok {
				delete(hub.Clients, client)
				close(client.Send)
			}
			hub.Mutex.Unlock()
		case message := <-hub.Broadcast:
			hub.Mutex.Lock()
			for client := range hub.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(hub.Clients, client)
				}
			}
			hub.Mutex.Unlock()
		}
	}
}

func (client *Client) ReadMessage(hub *Hub) {

	defer func() {
		hub.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading this message", err)
			break
		}
		hub.Broadcast <- message
	}

}

func (client *Client) WriteMessages() {

	defer client.Conn.Close()

	for message := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error sending this message", err)
			break
		}
	}

}

func ServeWs(hub *Hub, w http.ResponseWriter, request *http.Request) {

	conn, err := upgrader.Upgrade(w, request, nil)

	if err != nil {
		log.Println("Error during websocket upgrade", err)
		return
	}

	client := &Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	hub.Register <- client

	go client.WriteMessages()
	go client.ReadMessage(hub)

}
