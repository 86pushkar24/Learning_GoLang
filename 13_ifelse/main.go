package main
import (
	"fmt"
)
func main() {
	fmt.Println("If-Else in Go")
	// Example 1: Basic if-else statement
	loginCount := 5
	if loginCount > 10 { // brace opens in the same line
		fmt.Println("You have logged in more than 10 times.")
	} else if loginCount == 10 {
		fmt.Println("You have logged in exactly 10 times.")
	} else {
		fmt.Println("You have logged in less than or equal to 10 times.")
	}

	// Example 2: If statement with initialization
	if loginCount := 5; loginCount > 10 { // brace opens in the same line
		fmt.Println("You have logged in more than 10 times.")
	} else if loginCount == 10 {
		fmt.Println("You have logged in exactly 10 times.")
	} else {
		fmt.Println("You have logged in less than or equal to 10 times.")
	}
}