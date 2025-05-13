package main

import (
		"fmt"
	)	
func main() {
	fmt.Println("Maps in Go")

	// map declaration and initialization
	languages := make(map[string]string)
	languages["en"] = "English"
	languages["es"] = "Spanish"
	languages["fr"] = "French"

	fmt.Println("Languages:", languages) // map[en:English es:Spanish fr:French]
	fmt.Println("Language for 'en':", languages["en"]) // English

	delete(languages, "es") // delete Spanish

	// for range loop
	for key, value := range languages {
		fmt.Println("Key:", key, "Value:", value)
	}
	// second for range loop
	for _, value := range languages {
		fmt.Println("Value:", value)
	}
}