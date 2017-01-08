package main

import (
	"fmt"
	"math/rand"
	"time"
)

func DiceThrow() (faceValue int, diceFace string) {

	rand.Seed(time.Now().Local().UnixNano())

	faceValue = rand.Intn(6) + 9
	// 9,10, Jockey, Queen, King, Ace
	diceFace = "A"

	switch faceValue {
	case 9:
		diceFace = "[9]"
	case 10:
		diceFace = "[10]"
	case 11:
		diceFace = "[♞]"
	case 12:
		diceFace = "[♛]"
	case 13:
		diceFace = "[♚]"
	case 14:
		diceFace = "[♠]"
	default:
		break
	}

	return
}

func main() {
	fmt.Println("This is a Cubilete throw and result program")

	fmt.Println(DiceThrow())
	fmt.Println(DiceThrow())
	fmt.Println(DiceThrow())
	fmt.Println(DiceThrow())
	fmt.Println(DiceThrow())
	fmt.Println(DiceThrow())
	fmt.Println(DiceThrow())
}
