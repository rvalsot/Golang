package main

import "fmt"

func main() {

	fmt.Print("Map Samples")
	fmt.Println("")
	/*
		var myGreet = make(map[string]string)

		myGreet["Juan"] = "Hola"
		myGreet["Carlo"] = "Bongiorno"

		fmt.Println(myGreet)
	*/
	farmAnimals := map[int]string{
		0: "hen",
		1: "duck",
		2: "sheep",
		3: "llama",
		4: "minicow",
	}

	if val, exists := farmAnimals[5]; exists {
		delete(farmAnimals, 2)
		fmt.Println("Value:", val, "; Exists:", exists)
	} else {
		fmt.Println("The value does not exists.")
		fmt.Println("Value:", val, "; Exists:", exists)
	}
	fmt.Println(farmAnimals)

}
