//Darius Fiallo
//June 12, 2020
//TicTacToe
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var matrix = [3][3]string{
	{" ", " ", " "},
	{" ", " ", " "},
	{" ", " ", " "},
}
var game bool = false

func main() {

	game = false
	fmt.Println("Time for a game of tic tac toe. You are X, the computer is O.")
	count := 0
	for !game || count < 9 {
		if game == true {
			break
		}
		if count >= 9 {
			game = true
			fmt.Println("It's a tie.")
			break
		}
		yourMove()
		count += 2
	}

	fmt.Println("Final Score")

	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
	fmt.Println(matrix[2])

}

func yourMove() {

	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
	fmt.Println(matrix[2])

	fmt.Println("Here is the current board.")
	var user string = "X"

	var correct bool = false

	for !correct {
		fmt.Println("Now, top, middle, or bottom? 1/2/3")
		var row int
		fmt.Scanln(&row)

		fmt.Println("Now, left, middle, or right? 1/2/3")
		var col int
		fmt.Scanln(&col)

		if matrix[row-1][col-1] == " " {
			matrix[row-1][col-1] = user
			fmt.Println(user)
			correct = true
		} else {
			fmt.Println("Spot is filled. Try again")
			correct = false
		}
	}
	winCheck()
	if game == true {
		return
	}
	computerMove()
	winCheck()

}

func winCheck() {
	if matrix[0][0] == matrix[1][0] && matrix[1][0] == matrix[2][0] && (matrix[1][0] == "X" || matrix[1][0] == "O") {
		fmt.Println(matrix[0][0], " wins!")
		game = true
	} else if matrix[0][1] == matrix[1][1] && matrix[1][1] == matrix[2][1] && (matrix[1][1] == "X" || matrix[1][1] == "O") {
		fmt.Println(matrix[1][1], " wins!")
		game = true
	} else if matrix[0][2] == matrix[1][2] && matrix[1][2] == matrix[2][2] && (matrix[1][2] == "X" || matrix[1][2] == "O") {
		fmt.Println(matrix[1][2], " wins!")
		game = true
	} else if matrix[0][0] == matrix[0][1] && matrix[0][1] == matrix[0][2] && (matrix[0][1] == "X" || matrix[0][1] == "O") {
		fmt.Println(matrix[0][1], " wins!")
		game = true
	} else if matrix[1][0] == matrix[1][1] && matrix[1][1] == matrix[1][2] && (matrix[1][1] == "X" || matrix[1][1] == "O") {
		fmt.Println(matrix[1][1], " wins!")
		game = true
	} else if matrix[2][0] == matrix[2][1] && matrix[2][1] == matrix[2][2] && (matrix[2][1] == "X" || matrix[2][1] == "O") {
		fmt.Println(matrix[2][1], " wins!")
		game = true
	} else if matrix[0][0] == matrix[1][1] && matrix[1][1] == matrix[2][2] && (matrix[1][1] == "X" || matrix[1][1] == "O") {
		fmt.Println(matrix[1][1], " wins!")
		game = true
	} else if matrix[0][2] == matrix[1][1] && matrix[1][1] == matrix[2][0] && (matrix[1][1] == "X" || matrix[1][1] == "O") {
		fmt.Println(matrix[1][1], " wins!")
		game = true
	}
}

func computerMove() {
	var found bool = false

	if matrix[1][1] == "X" {
		switch {
		case matrix[0][0] == "X":
			if matrix[2][2] == "O" {
				break
			}
			matrix[2][2] = "O"
			found = true
			fmt.Println("Made it to bottom right.")
		case matrix[0][2] == "X":
			if matrix[2][0] == "O" {
				break
			}
			matrix[2][0] = "O"
			found = true
			fmt.Println("Made it to bottom left.")

		case matrix[2][0] == "X":
			if matrix[0][2] == "O" {
				break
			}
			matrix[0][2] = "O"
			found = true
			fmt.Println("Made it to top right corner.")
		case matrix[2][2] == "X":
			if matrix[0][0] == "O" {
				break
			}
			matrix[0][0] = "O"
			fmt.Println("Made it to top left corner.")
			found = true
		}
	}
	//Need to check that O's don't already exist there and to make a break statement like above if so
	if contains(matrix[0]) {
		switch {
		case matrix[0][0] == "X" && matrix[0][2] == "X":
			matrix[0][1] = "O"
			found = true
		case matrix[0][0] == "X" && matrix[0][1] == "X":
			matrix[0][2] = "O"
			found = true
		case matrix[0][1] == "X" && matrix[0][2] == "X":
			matrix[0][0] = "O"
			found = true
		}
	}
	if contains(matrix[1]) {
		switch {
		case matrix[1][1] == "X" && matrix[1][2] == "X":
			matrix[1][0] = "O"
			found = true
		case matrix[1][0] == "X" && matrix[1][1] == "X":
			matrix[1][2] = "O"
			found = true
		case matrix[1][0] == "X" && matrix[1][2] == "X":
			matrix[1][1] = "O"
			found = true
		}
	}
	if contains(matrix[2]) {
		switch {
		case matrix[2][1] == "X" && matrix[2][2] == "X":
			matrix[2][1] = "O"
			found = true
		case matrix[2][0] == "X" && matrix[2][1] == "X":
			matrix[2][2] = "O"
			found = true
		case matrix[2][1] == "X" && matrix[2][2] == "X":
			matrix[2][0] = "O"
			found = true
		}
	}

	rand.Seed(time.Now().UnixNano())
	for !found {
		i := rand.Intn(3)
		j := rand.Intn(3)

		if matrix[i][j] == " " {
			matrix[i][j] = "O"
			found = true
			break
		}

	}

}

func contains(row [3]string) bool {
	for i := 0; i < len(row); i++ {
		if row[i] == "X" {
			return true
		}
	}
	return false
}
