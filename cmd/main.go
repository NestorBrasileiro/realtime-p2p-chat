package main

import (
	"log"
	"net/http"

	"realtime-p2p-chat/internal/services"
)

func main() {

	hub := services.NewHub()
	go hub.Run()

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		services.ServeWs(hub, w, r)
	})

	log.Println("WebSocket Server running at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Error initializing the server: %v", err)
	}

}
