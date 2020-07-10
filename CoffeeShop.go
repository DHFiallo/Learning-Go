//Darius Fiallo
//June 4, 2020
//Coffee Shop

package main

import (
	"fmt"
	"strings"
)

func main() {
	var cost float64 = 0

	var size string
	var correct bool = true
	for correct == true {
		fmt.Println("Do you want small, medium, or large?")
		fmt.Scanln(&size)

		switch strings.ToUpper(size) {
		case "SMALL":
			cost += 2
			correct = false
		case "MEDIUM":
			cost += 3
			correct = false
		case "LARGE":
			cost += 4
			correct = false
		default:
			correct = true
			fmt.Println("Not a choice, choose again")

		}
	}

	var prepared string

	correct = true

	for correct == true {

		fmt.Println("Do you want brewed, espresso, or cold?")
		fmt.Scanln(&prepared)

		switch strings.ToUpper(prepared) {
		case "BREWED":
			cost += 0
			correct = false
		case "ESPRESSO":
			cost += .5
			correct = false
		case "COLD":
			cost += 1
			correct = false
		default:
			correct = true
			fmt.Println("Not a choice, choose again.")

		}
	}

	var flavoring string
	fmt.Println("Do you want flavored syrup? Yes/No?")
	fmt.Scanln(&flavoring)

	switch strings.ToUpper(flavoring) {
	case "YES":
		fmt.Println("What flavor? Hazelnut, vanilla, or caramel?")
		fmt.Scanln(&flavoring)
		cost += .5
	default:
		cost += 0
		flavoring = "no"
	}

	size = strings.ToLower(size)
	prepared = strings.ToLower(prepared)
	flavoring = strings.ToLower(flavoring)

	s := fmt.Sprintf("%.2f", cost)
	tip := fmt.Sprintf("%.2f", cost*1.15)

	fmt.Println("You asked for a " + size + " cup of " + prepared + " coffee with " + flavoring + " syrup.")
	fmt.Println("Your cup of coffee costs " + s)
	fmt.Println("The price with a tip is " + tip)

}
