package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	fmt.Println("JSON Encoder")
	fmt.Println("------------------")

	person1 := Person{"James", "Bond", 29, 007}
	json.NewEncoder(os.Stdout).Encode(person1)

	var personDecoded Person
	rdr := strings.NewReader(`{
		"First":"Hernan",
		"Last":"Cortes",
		"Age":1521
	}`)

	json.NewDecoder(rdr).Decode(&personDecoded)

	fmt.Println(personDecoded.Age)
	fmt.Println(personDecoded)

}
