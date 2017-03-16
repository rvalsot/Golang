package main

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
