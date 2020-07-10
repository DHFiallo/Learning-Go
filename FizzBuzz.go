//Darius Fiallo
//June 4th, 2020
//FizzBuzz

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var replaced int = 0

	var err error = nil
	var count string
	var count1 int
	for {
		fmt.Println("How many fizzing and buzzing units do you need in your life?")
		fmt.Scanln(&count)
		//Look at scanf for similar to java scanning

		count1, err = strconv.Atoi(count)
		fmt.Println(err)
		if err == nil {
			break
		}
	}

	if count1%1 == 0 {
		var i int = 0
		for {
			if replaced == count1 {
				fmt.Println("TRADITION!!")
				break
			}

			switch {
			case i == 0:
				fmt.Println("0")
				i++
				continue
			case i%5 == 0 && i%3 == 0:
				replaced++
				fmt.Println("fizz buzz")
			case i%3 == 0:
				replaced++
				fmt.Println("fizz")
			case i%5 == 0:
				replaced++
				fmt.Println("buzz")
			default:
				fmt.Println(strconv.Itoa(i))
			}
			i++
		}
	}
}
