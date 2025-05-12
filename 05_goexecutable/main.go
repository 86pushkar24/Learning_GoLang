package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	// Format the time
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)

	// Parse a string to time
	timeString := "2023-10-01 12:00:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed Time:", parsedTime)
}

// for building go executable for different OS and architecture
// first run go env to know the current OS and architecture
// GOOS="operating system" go build -o output_file_name main.go
// for example, to build for Windows 64-bit
// GOOS=windows GOARCH=amd64 go build -o hello.exe main.go
// to build for Linux 64-bit
// GOOS=linux GOARCH=amd64 go build -o hello main.go