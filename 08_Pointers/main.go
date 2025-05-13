package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers in Go")

	var ptr *int
	fmt.Println("Pointer:", ptr) // nil
	// fmt.Println("Value:", *ptr) // panic: runtime error: invalid memory address or nil pointer dereference
	// fmt.Println("Address:", &ptr) // panic: runtime error: invalid memory address or nil pointer dereference

	var number int = 42
	ptr = &number
	fmt.Println("Pointer:", ptr) // address of number - address of the variable
	fmt.Println("Value:", *ptr) // 42 - dereference the pointer
	fmt.Println("Address:", &ptr) // address of ptr - address of the pointer itself

	// operation on pointer
	*ptr = *ptr + 1
	fmt.Println("Value:", *ptr) // 43

}