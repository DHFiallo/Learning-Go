//Darius Fiallo
//June 3rd, 2020
//Lucky Seven game

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var balance int = 0
var highest int = 0
var total int = 0
var stopCount int = 0
var dice1 int = 0
var dice2 int = 0

func main() {

	fmt.Println("How many dollars you wanna bet?")
	fmt.Scanln(&balance)

	roll(dice1, dice2)

}
func rules(dice1, dice2 int) {
	fmt.Println("Dice1 ", strconv.Itoa(dice1), " Dice2 ", strconv.Itoa(dice2))
	if dice1+dice2 == 7 {
		balance += 4
	} else {
		balance -= 1
	}
	if balance > highest {
		stopCount = total
		highest = balance
	}
	if balance > 0 {
		roll(dice1, dice2)
	} else {
		fmt.Println("You are broke after ", strconv.Itoa(total), " rolls.")
		fmt.Println("You should have quit after ", strconv.Itoa(stopCount), " rolls when you had ", strconv.Itoa(highest))
	}
}

func roll(dice1, dice2 int) {
	total++
	rng := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(rng)
	dice1 = gen.Intn(6)
	dice2 = gen.Intn(6)
	rules(dice1, dice2)
}
