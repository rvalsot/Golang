package main

import "fmt"

const metersToFeet = 0.3048

func main() {
	var distance float64
	fmt.Println("Enter a meters distance:")
	fmt.Scan(&distance)
	feets := distance / metersToFeet
	fmt.Println(distance, " in meters, is ", feets, " in feets.")
}
