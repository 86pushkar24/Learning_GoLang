package main 
import "fmt"
func main() {
	// use fp as shortcut just like sysout in Java
	fmt.Println("Hello, Go!")
}


/*
	GOPATH defines your Go workspace and was essential for project and dependency management before Go modules. 
	For modern Go web backend development, prefer Go modules, but understanding GOPATH is useful when dealing with older code or specific environments
*/

// Environment Variables
// setting or piece of information-like a password, file path, or configuration-that programs can look up and use while they are running
// DB_URL to tell your web app where to find its database, or APP_MODE to switch between development and production settings