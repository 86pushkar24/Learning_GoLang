package main

import (
	"fmt"
	"strings"
	"os"
	"Bufio"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr) 								// Remove newline character
	num, err := strconv.Atoi(numStr) 								// Convert string to int
	if err != nil {													// Handle error
		fmt.Println("Error converting string to int:", err)
		return
	}
	fmt.Println("The number is:", num)
}