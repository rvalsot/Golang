package main

import "fmt"

func main() {
	var guess int
	option := 5

	fmt.Println("Please, give us a number:")
	fmt.Scan(&guess)

	if guess == option {
		fmt.Println("You've done it first time, its ", option)
	} else if guess > 5 {
		fmt.Println("Please, guess lower")
		fmt.Scan(&guess)

		if guess == option {
			fmt.Println("You've made it! It's", option)
		} else {
			fmt.Println("Sorry bro, try next time")
		}

	} else if guess < 5 {
		fmt.Println("Please, guess higher")
		fmt.Scan(&guess)

		if guess == option {
			fmt.Println("You've made it! It's:", option)
		} else {
			fmt.Println("Sorry bro, try next time")
		}
	}

}
