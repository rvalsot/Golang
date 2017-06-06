package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "Ya' sayin: %s \n", ln)
	}
	defer conn.Close()

	fmt.Println("Connection was shutdown")
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}

}
