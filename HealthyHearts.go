//Darius Fiallo
//June 2nd, 2020
//HealthyHearts.go

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Enter your age: ")
	var age string
	fmt.Scanln(&age)

	ageNum, _ := strconv.Atoi(age)

	fmt.Println("Your maximum heart rate should be " + strconv.Itoa((220 - ageNum)))

	var fiftyPercent = (0.5 * float32((220 - ageNum)))
	var eightyFivePercent = (0.85 * float32((220 - ageNum)))

	fmt.Println("Your target HR Zone is " + fmt.Sprintf("%f", fiftyPercent) + "-" + fmt.Sprintf("%f", eightyFivePercent) + " beats per minute.")

}
