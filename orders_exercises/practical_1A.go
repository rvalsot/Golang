/*
 * Location: 1525
 * Exercise Practical A
 * -------------------------
 * Step 1: Set up a Ticking Midmarket Spot
 * User inputs → 1.- Initial spot; 2.- Time between ticks; 3.- Variation per tick
 * Outputs: 1.- Current time step; 2.- Current spot
 * Pause/Continue buttons are required
 */

package main

import "fmt"

func main() {
	fmt.Println("Practical A, Task A")

	// User inputs declaration
	var initialSpot float64
	var timeBetweenTicks int
	var variationPerTick float64

	fmt.Println("Please provide:")
	fmt.Println("Initial Spot price")
	fmt.Scan(&initialSpot)
	fmt.Println("Time between ticks")
	fmt.Scan(&timeBetweenTicks)
	fmt.Println("Variation per tick")
	fmt.Scan(&variationPerTick)

	fmt.Println("1.-", initialSpot, "2.-", timeBetweenTicks, "3.-", variationPerTick)

}
