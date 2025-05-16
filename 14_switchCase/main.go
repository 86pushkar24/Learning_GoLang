package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch Case in Go")

	day := rand.Intn(7) + 1 // Random number between 1 and 7
	switch day {
	case 1: 
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
		fallthrough // fallthrough allows execution to continue to the next case
		// This is useful if you want to execute the next case's code as well
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}