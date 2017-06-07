package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "Hi from a TCP server")
		fmt.Fprintln(conn, "How you doin?")
		fmt.Fprintf(conn, "%v", "Well, I hope so!")
		conn.Close()
	}

}
