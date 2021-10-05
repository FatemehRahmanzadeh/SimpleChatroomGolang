package main

import (
	"io"
	"log"
	"net/http"

	"github.com/FatemehRahmanzadeh/chat_sample/auth"
	"github.com/FatemehRahmanzadeh/chat_sample/config"
	"github.com/FatemehRahmanzadeh/chat_sample/repository"
	"github.com/gorilla/websocket"
)

type ChatMessage struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan ChatMessage)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()
	clients[ws] = true

	for {
		var msg ChatMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		// send new message to the channel
		broadcaster <- msg
	}
}

// If a message is sent while a client is closing, ignore the error
func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func handleMessages() {
	for {
		// grab any next message from channel
		msg := <-broadcaster
		messageClients(msg)
	}
}

func messageClients(msg ChatMessage) {
	// send to every client currently connected
	for client := range clients {
		messageClient(client, msg)
	}
}

func messageClient(client *websocket.Conn, msg ChatMessage) {
	err := client.WriteJSON(msg)
	if err != nil && unsafeError(err) {
		log.Printf("error: %v", err)
		client.Close()
		delete(clients, client)
	}
}

func main() {
	db := config.InitDB()
	userRepository := &repository.UserRepository{Db: db}
	api := &API{UserRepository: userRepository}

	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/websocket", auth.AuthMiddleware(handleConnections))
	// http.HandleFunc("/websocket", handleConnections)
	go handleMessages()
	http.HandleFunc("/api/login", api.HandleLogin)

	log.Print("Server starting at localhost:4444")

	if err := http.ListenAndServe(":4444", nil); err != nil {
		log.Fatal(err)
	}
}
