package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Initialization ______________________________________________________________

var upgrader = websocket.Upgrader{}

// Handlers ____________________________________________________________________
func socketHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.gohtml")
}

func socketV1(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error at conn := upgrader.Upgrade: ", err)
	}

	// Goroutine to read messages from client
	go func(conn *websocket.Conn) {
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("ReadMessage error: \n", err)
			}

			// Return message creation
			returnClause := "Returned message: " + string(msg)
			returnMsg := []byte(returnClause)
			// We will return to his guy his message
			conn.WriteMessage(msgType, returnMsg)
		}
	}(conn)

}

func main() {
	http.HandleFunc("/", socketHandler)
	http.HandleFunc("/v1/ws", socketV1)

	port := ":8070"

	fmt.Println("Listening and Serving at port", port)
	http.ListenAndServe(port, nil)
}
