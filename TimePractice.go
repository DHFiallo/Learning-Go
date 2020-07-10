//Darius Fiallo
//June 24, 2020
//Time Practice

package main

import (
	"fmt"
	"time"
)

func main() {
	// display current time
	now := time.Now()
	fmt.Println("Today's date and time:", now)
	fmt.Println("Current year:", now.Year())
	fmt.Println("Current month:", now.Month())
	fmt.Println("Current day:", now.Day())
	fmt.Println("Current hour:", now.Hour())
	fmt.Println("Current minute:", now.Minute())
	fmt.Println("Current second:", now.Second())
	fmt.Println("Current nanosecond:", now.Nanosecond())
	fmt.Println("Current location:", now.Location())
	fmt.Println("Current weekday:", now.Weekday())

	// create custom time
	customTime := time.Date(
		2025, 05, 15, 15, 20, 00, 0, time.Local)
	fmt.Println("Custom date and time:", customTime)

	fmt.Println("Custom year:", customTime.Year())
	fmt.Println("Custom month:", customTime.Month())
	fmt.Println("Custom day:", customTime.Day())
	fmt.Println("Custom weekday:", customTime.Weekday())
	fmt.Println("Custom hour:", customTime.Hour())
	fmt.Println("Custom minute:", customTime.Minute())
	fmt.Println("Custom second:", customTime.Second())
	fmt.Println("Custom nanosecond:", customTime.Nanosecond())
	fmt.Println("Custom location:", customTime.Location())
}
