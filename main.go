package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// bufio.NewReader(os.Stdin) creates a new reader that reads from standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	// This is comma , ok idiom in Go
	// The first value is the input string, and the second value is an error (if any)
	// If you don't want to check for errors, you can ignore the second value using _
	name, _ := reader.ReadString('\n')
	// ReadString('\n') reads until the newline character

	// Print the name entered by the user
	fmt.Println("Hello,", name)
	fmt.Printf("The type of variable is %T\n", name)
}