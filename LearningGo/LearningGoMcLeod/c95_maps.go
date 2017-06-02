package main

import "fmt"

func main() {
	fmt.Print("Maps")
	fmt.Println("")

	m := make(map[string]int)

	m["K1"] = 7
	m["K2"] = 12
	m["K4"] = 13

	fmt.Println("Map:", m)

	v1 := m["K1"]
	fmt.Printf("%T \n", v1)
	fmt.Println("m map length:", len(m))
	delete(m, "K2")
	fmt.Println("After deleting K2: ", m)

	_, prs := m["K2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Map n:", n)
}
