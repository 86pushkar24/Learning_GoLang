package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays in Go")

	// array declaration
	var fruitList [5]string

	// array initialization
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[4] = "Date"

	fmt.Println("Fruit List:", fruitList) // [Apple Banana   Date]
	fmt.Println("Fruit List Length:", len(fruitList)) // 5

	// array declaration and initialization
	vegetableList := [5]string{"Carrot", "Potato", "Tomato", "Onion", "Cabbage"}
	fmt.Println("Vegetable List:", vegetableList) // [Carrot Potato Tomato Onion Cabbage]
	fmt.Println("Vegetable List Length:", len(vegetableList)) // 5
}