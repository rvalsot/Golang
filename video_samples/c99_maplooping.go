package main

import "fmt"

func main() {
	fmt.Print("Map Looping")
	fmt.Println(" ")

	preciousElements := map[string]string{
		"Au": "Gold",
		"Ag": "Silver",
		"Pt": "Platinum",
		"Al": "Aluminium",
		"Rh": "Rhodium",
	}

	for key, val := range preciousElements {
		fmt.Println(key, "-", val)
	}

}
