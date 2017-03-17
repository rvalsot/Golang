package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	// forward is a channel that holds incoming messages that should be forwarded to the other clients
	forward chan []byte

	// Join is a channel for clients wishing to join the room.
	join chan *client

	// Leave is a channel for clients wishing to leave the room .
	leave chan *client

	// Clients hoolds all current clients in this room.
	clients map[*client]bool
}

// New room constructor that will be ready to roll
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

// Concurrency: one at the time entry, leave, and delivery
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining client
			r.clients[client] = true
		case client := <-r.leave:
			// leaving client
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the message
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// Socket creation
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// ServeHTTP method to be used by http.Handle function
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP at Room Fail:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
