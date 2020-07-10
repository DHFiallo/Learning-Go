package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("\"Hello, world!\"")
	fmt.Print("Go is fun!\nGo is also easy.\n")

	fmt.Println("What's your name?")
	var name string
	fmt.Scanln(&name)

	fmt.Println("In order to play, select a choice " + name + "\n1. Easy Mode: 1-5,\n2. Normal Mode: 1-20,\n3. Hard Mode:1-50\n (Input a number)")
	var choice string
	fmt.Scanln(&choice)

	var num int
	var loop int

	switch choice {
	case "1":
		num = rand.Intn(5)
		loop = 5
	case "2":
		num = rand.Intn(20)
		loop = 20
	case "3":
		num = rand.Intn(50)
		loop = 50
	default:
		fmt.Println("Wrong choice")
	}

	var attempts int = 0

	for i := 0; i < loop; i++ {
		fmt.Println("Guess a number, ", name)
		var guess string
		fmt.Scanln(&guess)
		attempts++
		var guessInt, _ = strconv.Atoi(guess)
		if guessInt > loop {
			fmt.Println("You entered an invalid number, one that is outside the range.")
		}
		if (guessInt == num) && (attempts == 0) {
			fmt.Println("Hot diggity dog, you did it in 1 try, ", name, "!")
			break
		}
		if guessInt == num {
			fmt.Println("Nice! It took you ", attempts, " attempt(s), ", name)
			break
		}

	}

	/*if cat == "Y" {
		fmt.Print("You got any dogs? Y/N")
		var dog string
		fmt.Scanln(&dog)
		if dog == "Y" {
			fmt.Println("You must really love pets!")
		} else {
			fmt.Println("That's cool, cat lady.")
		}
	} else {
		fmt.Println("Maybe you need more pets.")
	}*/

	/*fmt.Print("Enter your letter grade")
	var grade string
	fmt.Scanln(&grade)

	switch grade {
	case "A":
		fmt.Println("You got an A.")
	case "B":
		fmt.Println("B, pretty good.")
	case "C":
		fmt.Println("C's get degrees.")
	case "D":
		fmt.Println("Got a D, yay..")
	default:
		fmt.Println("Wow, you fucking suck.")
	}
	*/

}
