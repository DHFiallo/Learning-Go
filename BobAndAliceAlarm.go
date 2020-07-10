//Darius Fiallo, Jasmin Smith, Andrew Horn
//June 19, 2020
//Bob and Alice Alarm

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//func (function name) (variable name chan variable type)

func bobMorningRoutine(limit int) {

	fmt.Println("Bob started getting ready.")

	time.Sleep(time.Duration(limit) * time.Second)

	fmt.Println("Bob spent " + strconv.Itoa(limit) + " seconds getting ready.")
}

func aliceMorningRoutine(limit int) {

	fmt.Println("Alice started getting ready.")

	time.Sleep(time.Duration(limit) * time.Second)

	fmt.Println("Alice spent " + strconv.Itoa(limit) + " seconds getting ready.")
}

func aliceShoes(limit int) {

	fmt.Println("Alice is putting on shoes")

	time.Sleep(time.Duration(limit) * time.Second)

	fmt.Println("Alice spent " + strconv.Itoa(limit) + " seconds putting on shoes")

}

func bobShoes(limit int) {

	fmt.Println("Bob is putting on shoes")

	time.Sleep(time.Duration(limit) * time.Second)

	fmt.Println("Bob spent " + strconv.Itoa(limit) + " seconds putting on shoes")

}

func main() {

	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)

	bob := gen.Intn(31) + 60
	alice := gen.Intn(31) + 60
	// b := rand.Intn(31) + 60
	// a := rand.Intn(31) + 60
	go bobMorningRoutine(bob)
	go aliceMorningRoutine(alice)

	var max int
	if bob > alice {
		max = bob
	} else {
		max = alice
	}
	time.Sleep(time.Duration(max) * time.Second) //Set it equal to greater of 2 integers
	fmt.Println("Arming alarm.")
	bob = gen.Intn(11) + 35
	alice = gen.Intn(11) + 35
	if bob > alice {
		max = bob
	} else {
		max = alice
	}
	go bobShoes(bob)
	go aliceShoes(alice)
	fmt.Println("Alarm is counting down.")
	time.Sleep(time.Duration(max) * time.Second)
	fmt.Println("Exiting and locking the door.")
	time.Sleep(time.Duration(60-max) * time.Second)
	fmt.Println("Alarm is armed.")

	//for the time between exiting and locking the door and alarm is armed
}
