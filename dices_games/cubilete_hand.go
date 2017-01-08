package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*---------------------------------------------------------------------------*/

func CubileteHand() (cubileteHand [5]int) {
	rand.Seed(time.Now().Local().UnixNano())

	for i := 0; i < 5; i++ {
		auxiliar := rand.Intn(6) + 9
		cubileteHand[i] = auxiliar
	}

	return
}

/*---------------------------------------------------------------------------*/

func CubileteHandStringer(numericHand [5]int) (stringHand [5]string) {
	selector := func(value int) (icon string) {
		switch value {
		case 9:
			icon = "9"
		case 10:
			icon = "10"
		case 11:
			icon = "♞"
		case 12:
			icon = "♛"
		case 13:
			icon = "♚"
		case 14:
			icon = "♠"
		default:
			break
		}
		return icon
	}

	for i := 0; i < 5; i++ {
		stringHand[i] = selector(numericHand[i])
	}

	return
}

/*---------------------------------------------------------------------------*/

func CubileteResult(numericHand [5]int) map[int]int {

	resultsMap := make(map[int]int, 5)

	for _, item := range numericHand {
		_, exist := resultsMap[item]

		if exist {
			resultsMap[item]++
		} else {
			resultsMap[item] = 1
		}
	}

	return resultsMap
}

func StringifyResults(results map[int]int) map[string]int {
	stringMap := make(map[string]int, 5)

	for key, value := range results {
		fmt.Printf("Key: %d, Value: %d", key, value)
	}
}

/*---------------------------------------------------------------------------*/

func main() {
	fmt.Println("This is a full hand show of cubilete")

	myHand := CubileteHand()
	fmt.Println("My hand's numbers:", myHand)

	fmt.Println(CubileteResult(myHand))

	// myFaces := CubileteHandStringer(myHand)
	// fmt.Println(myFaces)

}
