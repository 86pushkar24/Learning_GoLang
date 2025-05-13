package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in Go")
	// Similar to classes in OOP
	// no inheritance
	// no super or base class or parent class

	// struct declaration // Person P is capitalized
	type Person struct {
		Name   string
		Age    int
		Email  string
		isLearning bool
	}

	// struct initialization
	p1 := Person{
		Name:   "John Doe",
		Age:    30,
		Email:  "john.doe@example.com",
		isLearning: true,
	}
	p2 := Person{"Jane Doe", 25, "jane.doe@example.com", false}

	fmt.Println("Person 1:", p1) // Person 1: {John Doe 30 john.doe@example.com true}
	fmt.Printf("Person 2: %+v\n", p2) // Person 2: {Name:Jane Doe Age:25 Email:jane.doe@example.com isLearning:false}
	fmt.Println("Name:", p1.Name) // Name: John Doe
	fmt.Printf("Age: %d and isLearning: %t\n", p1.Age, p1.isLearning) // Age: 30 and isLearning: true
}