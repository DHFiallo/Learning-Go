// Darius Fiallo
// June 1st, 2020
// This program prints off some information about myself
package main

import "fmt"

func main() {
	var custName string

	fmt.Print("Enter your first name: ") //Prompts user
	var firstName string
	fmt.Scanln(&firstName)

	fmt.Print("Enter your last name: ") //Prompts user
	var lastName string
	fmt.Scanln(&lastName)

	custName = firstName + " " + lastName

	fmt.Println("So your name is ", custName)

	fmt.Print("Enter your house number: ") //Prompts user
	var houseNum string
	fmt.Scanln(&houseNum)

	fmt.Print("Enter your street name: ") //Prompts user
	var streetName string
	fmt.Scanln(&streetName)

	fmt.Print("Enter your city: ") //Prompts user
	var city string
	fmt.Scanln(&city)

	fmt.Print("Enter your state abbreviation: ") //Prompts user
	var state string
	fmt.Scanln(&state)

	fmt.Print("Enter your zipcode: ") //Prompts user
	var zipcode string
	fmt.Scanln(&zipcode)

	var fullAddress = houseNum + " " + streetName + " " + city + " " + state + " " + zipcode
	fmt.Println("So your full address is ", fullAddress)

	var mainGreeting string = "Hello world!"
	var year string = "The year is 2020"
	fmt.Println(mainGreeting, " ", year)

	fmt.Println("Accessing the address of a variable")
	fmt.Println(&custName)
}

/*
This program is a modified verison of Hello World
but it introduces me instead.
*/
