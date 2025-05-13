package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")
	
	// slice declaration and initialization
	fruitList := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	// slice append
	fruitList = append(fruitList, "Fig", "Guava")
	fmt.Println("Fruit List:", fruitList) // [Apple Banana Cherry Date Elderberry Fig Guava]
	fmt.Println(fruitList[1:]);
	fmt.Println(fruitList[:5]); // 5 is exclusive
	fmt.Println(fruitList[1:3]);

	// slice length
	fmt.Println("Fruit List Length:", len(fruitList)) // 7
	// slice capacity
	fmt.Println("Fruit List Capacity:", cap(fruitList)) // 10

	numList := make([]int, 5)  
	numList[0] = 5
	numList[1] = 4
	numList[2] = 3
	numList[3] = 2
	numList[4] = 1
	// numList[5] = 5 // This will cause an error: index out of range
	// Append does not cause an error because it reallocates the memory
	numList = append(numList, 6, 7)
	sort.IntsAreSorted(numList) 
	sort.Ints(numList)

	// remove an element from a slice -- super important
	numList = append(numList[:2], numList[3:]...) // remove element at index 2
	fmt.Println("Removed Element:", numList) // [5 4 2 3 6 7]
}