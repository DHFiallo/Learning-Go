//Darius Fiallo
//June 2nd, 2020
//DogBreed.go

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	fmt.Print("What's the name of your dog?")
	var dogName string
	fmt.Scanln(&dogName)

	ns := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(ns)

	var first = randGen.Intn(50)
	var second = randGen.Intn(75 - first)
	var third = randGen.Intn(80 - (first + second))
	var fourth = randGen.Intn(100 - (first + second + third))
	var fifth = 100 - (first + second + third + fourth)

	fmt.Println(dogName + " is:")

	fmt.Println(strconv.Itoa(first) + "% St. Bernard")
	fmt.Println(strconv.Itoa(second) + "% Chihuaha")
	fmt.Println(strconv.Itoa(third) + "% Doberman")
	fmt.Println(strconv.Itoa(fourth) + "% Poodle")
	fmt.Println(strconv.Itoa(fifth) + "% Maltese")

}
