package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Via a pseudo random number from 1-6, determines which face of a dice to return
func DiceRoll() (number int, diceFace string) {

	// Seeding
	rand.Seed(time.Now().UTC().UnixNano())

	diceFace = "1"

	number = rand.Intn(6) + 1

	switch number {
	case 1:
		diceFace = "⚀"
	case 2:
		diceFace = "⚁"
	case 3:
		diceFace = "⚂"
	case 4:
		diceFace = "⚃"
	case 5:
		diceFace = "⚄"
	case 6:
		diceFace = "⚅"
	default:
		break

	}

	// fmt.Println(number)

	return
}

func main() {
	fmt.Println("This is a dice roller")

	fmt.Println(DiceRoll())
}
