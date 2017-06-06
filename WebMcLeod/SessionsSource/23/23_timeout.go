package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(15 * time.Second))
	if err != nil {
		log.Println("Connection timeout")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "You say: %s", ln)
	}
	defer conn.Close()

	fmt.Println("Connection terminated")

}

func main() {
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}
