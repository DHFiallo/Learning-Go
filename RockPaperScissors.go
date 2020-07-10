package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var user int = 0
var computer int = 0
var ties int = 0

func main() {
	var rounds int = 0
	for rounds > 10 || rounds < 1 {
		fmt.Println("How many rounds do you want to play? 1-10")
		fmt.Scanln(&rounds)
		if rounds > 10 || rounds < 1 {
			fmt.Println("Wrong number of rounds")
		} else {
			game(rounds)
		}
	}
}
func userChoice() int {
	var userchoiceInt int = 0
	for {
		fmt.Println("Please enter your choice:\nRock - 1\nPaper - 2\nScissors -3")
		fmt.Scanln(&userchoiceInt)
		//loop will continue until player makes valid choice
		if userchoiceInt > 0 && userchoiceInt < 4 {
			break
		}
	}
	return userchoiceInt
}

//function takes two guesses and increments the global variable based on the result.
func compare(userGuess int, computerGuess int) {
	//1 = rock 2 = paper 3 = scissors. Game follows traditional rules
	if userGuess == 1 && computerGuess == 1 {
		ties++
		fmt.Println("It's a tie")
	}
	if userGuess == 1 && computerGuess == 2 {
		computer++
		fmt.Println("I win")
	}
	if userGuess == 1 && computerGuess == 3 {
		user++
		fmt.Println("You win")
	}
	if userGuess == 2 && computerGuess == 1 {
		user++
		fmt.Println("You win")
	}
	if userGuess == 2 && computerGuess == 2 {
		ties++
		fmt.Println("It's a tie")
	}
	if userGuess == 2 && computerGuess == 3 {
		computer++
		fmt.Println("I win")
	}
	if userGuess == 3 && computerGuess == 1 {
		computer++
		fmt.Println("I win")
	}
	if userGuess == 3 && computerGuess == 2 {
		user++
		fmt.Println("You win")
	}
	if userGuess == 3 && computerGuess == 3 {
		ties++
		fmt.Println("It's a tie")
	}
}

//Generates a rock, paper, or scissors for the computer
func computerChoice() int {
	ns := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(ns)

	var compChoice int = gen.Intn(3) + 1

	return compChoice
}

//Calls and starts the entire game
func game(rounds int) {

	user = 0
	computer = 0
	ties = 0

	for i := 1; i <= rounds; i++ {
		comp := computerChoice()
		user := userChoice()
		compare(user, comp)
	}
	fmt.Println("User Wins: ", strconv.Itoa(user))
	fmt.Println("Computer Wins: ", strconv.Itoa(computer))
	fmt.Println("Ties: ", strconv.Itoa(ties))

	if user > computer {
		fmt.Println("User wins!")
	} else if computer > user {
		fmt.Println("Computer wins.")
	} else {
		fmt.Println("Tied!")
	}

	var choice string
	fmt.Println("Want to play again? YES/NO")
	fmt.Scanln(&choice)

	if strings.ToUpper(choice) == "YES" {
		fmt.Println("How many rounds?")
		fmt.Scanln(&rounds)

		if rounds > 10 || rounds < 1 {
			fmt.Println("Error!")
			os.Exit(1)
		}
		game(rounds)
	} else {
		fmt.Println("Thanks for playing!")
	}
}
