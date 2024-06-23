package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

type Message struct {
	Action  string  `json:"action"`
	Time    float64 `json:"time"`
	Content string  `json:"content,omitempty"`
	Sender  string  `json:"sender,omitempty"`
}

func main() {
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/ws", handleConnections)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	go handleMessages()

	log.Println("HTTPS server started on :8080")
	//err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Check the URL and set the sender's name accordingly
		if r.Host == "localhost:8080" {
			msg.Sender = "Vinay"
		} else {
			msg.Sender = "other"
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
