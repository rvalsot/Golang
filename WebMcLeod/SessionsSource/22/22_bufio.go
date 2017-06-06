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
	}
	defer conn.Close()

	// How does the above reader know when it's done?
	fmt.Println("Code got here")
}

func main() {
	listen, err := net.Listen("tcp", ":8020")
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}
