package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	fmt.Println("JSON Marshall")
	fmt.Println("------------------")

	person1 := Person{"James", "Bond", 29, 007}
	bs, _ := json.Marshal(person1)

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

	var byteperson Person
	fmt.Print(byteperson)

	bs2 := []byte(`{
			"First" : "Miguel",
			"Last"	: "Hidalgo",
			"Age"		:	20,
			"notExported" : 1820
		}`)

	json.Unmarshal(bs2, &byteperson)
	fmt.Println("-------------------------")
	fmt.Printf("%T", byteperson)
	fmt.Println(byteperson)
}
