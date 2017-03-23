package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hi!")

	router := NewRouter()
	fmt.Println("New Router created")
	log.Fatal(http.ListenAndServe(":8080", router))
}
