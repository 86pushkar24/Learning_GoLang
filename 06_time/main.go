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
	formattedTime := currentTime.Format("2006-01-02 15:04:05 Monday")
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

// default time format
// 2006-01-02 15:04:05.999999999 Monday -0700 MST

