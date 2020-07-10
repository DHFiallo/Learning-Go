//Darius Fiallo
//June 2nd, 2020
//Flow of Control Exercises

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	/*var incorrect int8 = 0
	var correct int8 = 0
	var userAnswer string
	var questions int8

	for i := 0; i < 10; i++ {
		var one, two int
		one = rand.Intn(99)
		two = rand.Intn(99)
		fmt.Println("What is ", strconv.Itoa(one), " + ", strconv.Itoa(two), "?")
		fmt.Scanln(&userAnswer)
		questions++

		var answer int = (one + two)
		if userAnswer == strconv.Itoa(answer) {
			correct++
			incorrect = 0
			percentage := fmt.Sprintf("%.2f", (float32(correct)/float32(questions))*100)
			fmt.Println("Correct!")
			fmt.Println("You've answered ", questions, " questions so far")
			fmt.Println("Your overall accuracy is " + percentage)
		} else {
			incorrect++
			percentage := fmt.Sprintf("%.2f", float32(correct)/float32(questions)*100)
			fmt.Println("Incorrect!")
			fmt.Println("You've answered ", questions, " questions so far")
			fmt.Println("Your overall accuracy is " + percentage)
		}

		if incorrect == 3 {
			fmt.Println("You lost. You got 3 incorrect")
		}

	}*/

	/*fmt.Println("Please give a url")
	var url string
	fmt.Scanln(&url)
	var first string
	var last string = string(url[len(url)-1]) //Grabs the last letter
	var num int = rand.Intn(99)               //Generates the 2 digit random number

	var short string = "http://bigTiny.com/" //Sets the base of the shortened url

	words := strings.Split(url, "/") //Splits the user url apart
	word := words[3]                 //Grabs the word after .com
	first = string(word[0])          //Grabs the first letter of that word

	short += first + last + strconv.Itoa(num)
	fmt.Println(short)*/

	/*phoneNum := regexp.MustCompile("\\(?[2-9]{3}\\)?-?[0-9]{3}-?[0-9]{4}")

	split := regexp.MustCompile("[0-9]{3,4}?")
	fmt.Println("Please input a phone number")
	var phone string
	fmt.Scanln(&phone)

	splitPhone := split.FindAllString(phone, -1)
	var formattedPhone string = strings.Join(splitPhone, "-") + string(phone[len(phone)-1])

	if phoneNum.MatchString(phone) {
		fmt.Println("That is, indeed, a valid phone number. Here is " +
			"the formatted version " + formattedPhone)
	}*/

	/*fmt.Println("Enter an email address:")
	var email string
	fmt.Scanln(&email)

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if emailRegex.MatchString(email) {
		fmt.Println("This is valid.")
	} else {

		fmt.Println("This isn't valid")
	}*/

	var user1 int8
	var user2 int8
	var dealer1 int8
	var dealer2 int8

	dealer1 = int8(rand.Intn(10))
	user1 = int8(rand.Intn(10))
	fmt.Println("Your first card is: ", strconv.Itoa(int(user1)))

	dealer2 = int8(rand.Intn(10))
	fmt.Println("My card is: ", strconv.Itoa(int(dealer2)))

	user2 = int8(rand.Intn(10))
	fmt.Println("Your second card is: ", strconv.Itoa(int(user2)))

	fmt.Println("Hit or stay? Type 1 for Hit, 2 for Stay")
	var choice string
	fmt.Scanln(&choice)

	if choice == "1" {
		user3 := int8(rand.Intn(10))
		fmt.Println("Your third card is: ", strconv.Itoa(int(user3)))
		if((user1+user2+user3) == 21)
		{
			fmt.Println("Bust!")
		} else if ((user1+user2+user3) > 21)) {
			fmt.Println("My cards are " + strconv.Itoa(int(dealer1)) + " and " + strconv.Itoa(int(dealer2)))
			if()
		}
	}

}
