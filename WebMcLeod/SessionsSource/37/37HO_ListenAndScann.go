package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func serve(c net.Conn) {
	defer c.Close()

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln)
		if ln == "" {
			fmt.Println("This is the end of the HTTP request headers")
			break
		}
	}
}

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
		go serve(conn)
	}

}
