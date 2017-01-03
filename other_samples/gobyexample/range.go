package main

import "fmt"

func main() {
	numeros := []int{2, 3, 4, 5}
	suma := 0
	for _, numero := range numeros {
		suma += numero
		fmt.Println("Suma:", suma)
	}
	for i, numero := range numeros {
		if numero == 3 {
			fmt.Println("Indice:", i)
		}
	}

	fruitMap := map[string]string{
		"a": "apple",
		"b": "banana",
		"m": "mango",
		"o": "orange",
		"p": "pear",
	}

	for key, value := range fruitMap {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range fruitMap {
		fmt.Println("Key:", key)
	}

	for i, c := range "golang!" {
		fmt.Println(i, c)
	}

}
