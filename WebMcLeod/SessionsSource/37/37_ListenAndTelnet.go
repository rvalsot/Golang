package main

import (
	"io"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8079")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "TCP connection konws you're here")
		conn.Close()
	}

}
