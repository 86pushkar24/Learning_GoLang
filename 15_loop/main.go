package main
import (
	"fmt"
)

func main() {
	fmt.Println("For Loop in Go")
	
	// Example 1: Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
	
	// Example 2: For loop with a condition
	j := 0
	for j < 5 {
		fmt.Println("Iteration:", j)
		j++
	}
	
	// Example 3: Infinite for loop (use with caution)
	k := 0
	for {
		if k >= 5 {
			break // Break out of the loop when k is greater than or equal to 5
		}
		fmt.Println("Iteration:", k)
		k++
	}
}