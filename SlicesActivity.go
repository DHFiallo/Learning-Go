//Darius Fiallo
//June 13th, 2020
//Slices Activity

package main

import "fmt"

func main() {
	var array [10]string

	var total int
	for i := 0; i < len(array); i++ {
		var user string
		fmt.Println("Give me a word.")
		fmt.Scanln(&user)

		total += len(user)

		array[i] = user
	}
	var greater []string //slice for words greater than average

	avg := total / 10
	fmt.Println("The average")
	fmt.Println(avg)

	for i := 0; i < len(array); i++ {
		if len(array[i]) > avg {
			greater = append(greater, array[i])
		}
	}
	var lesser []string

	for i := 0; i < len(array); i++ {
		if len(array[i]) < avg {
			lesser = append(lesser, array[i])
		}
	}

	fmt.Println(array)
	fmt.Println("The greater")
	fmt.Println(greater)
	fmt.Println("The lesser")
	fmt.Println(lesser)
}
