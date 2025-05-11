package main

import "fmt"

// constants -- initialized with Capital letter and have PUBLIC Access
const LoginToken string = "gibrishh" // constant value of PI

func main() {
	var userName string = "Pushkar"
	fmt.Println("Hello, " + userName + "!") 
	fmt.Printf("The type of variable is %T\n", userName) 

	var age int = 25
	fmt.Println("Your age is", age)
	fmt.Printf("The type of variable is %T\n", age)

	var isStudent bool = true
	fmt.Println("Are you a student?", isStudent)
	fmt.Printf("The type of variable is %T\n", isStudent)

	// Implicit Declaration
	var coordinates = 3.14
	fmt.Println("Coordinates:", coordinates)
	fmt.Printf("The type of variable is %T\n", coordinates)

	// Short Declaration
	// := , only inside functions, not at the package level , infers type from vaue at rhs
	city := "New York"
	fmt.Println("City:", city)
	fmt.Printf("The type of variable is %T\n", city)

	fmt.Println("Login Token:", LoginToken)
	fmt.Printf("The type of variable is %T\n", LoginToken)
}