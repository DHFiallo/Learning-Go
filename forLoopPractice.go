//Darius Fiallo
//June 3rd, 2020
//Practice Activities: For Loops

package main

func main() {

	//Prints out the alphabet
	/*abc := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 1; i <= len(abc); i++ {
		fmt.Println(abc[i-1 : i])
	}*/

	//Displays even numbers
	/*for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("Even number: " + strconv.Itoa(i))
		}
	}*/

	//Displays odd numbers
	/*for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println("Odd number: " + strconv.Itoa(i))
		}
	}*/

	//Finds prime numbers between 0 and 100. O(n^2)
	/*for i := 0; i <= 100; i++ {
		prime := true

		for j := 2; j <= i; j++ {

			if i == j {
				continue
			}

			var x = float64(i) / float64(j)
			if math.Mod(x, 1.0) == 0 { // Or could just do mod of i and j and see if equal to 0
				prime = false
				break
			} else {
				prime = true
			}
		}
		if prime == true { //Or just if prime
			fmt.Println(strconv.Itoa(i) + " is prime.")
		}

	}*/

	//Sum of numbers from 0 to 100
	/*total := 0
	for i := 0; i <= 100; i++ {
		total += i
	}
	fmt.Println(strconv.Itoa(total))*/

	//Prints every number divisible by 50 between 100 and 1000
	/*for i := 100; i <= 1000; i += 50 {
		fmt.Println(strconv.Itoa(i))
	}*/
	//Can also do this with range function, although need to know how to make arrays

	/*Prints out answer for exponent problem
	var x int
	var y int
	fmt.Println("We're doing exponents. x^y. Give an x value")
	fmt.Scanln(&x)

	fmt.Println("Give a y value")
	fmt.Scanln(&y)

	total := 1
	for i := 0; i < y; i++ {
		total = total * x
	}
	fmt.Println(strconv.Itoa(total))*/

	/* Prints a multiplication table
	var x int
	fmt.Println("We're doing a multiplication table. Enter a digit")
	fmt.Scanln(&x)

	for i := 1; i <= x; i++ {
		fmt.Println(strconv.Itoa(i) + " times " + strconv.Itoa(x) + " is " + strconv.Itoa(i*x))
	}*/

	/*Activity 8, prints out sum of numbers up to that digit
	var x int
	fmt.Println("Activity 8. Enter a digit")
	fmt.Scanln(&x)

	var total int

	for i := 1; i <= x; i++ {
		total += i
		fmt.Println("Total: " + strconv.Itoa(total))

	}*/

	//Activity 16
	/*letter := "Hello world"
	count := 0
	for i, s := range letter {
		fmt.Println(i, string(s))
		count++
	}
	fmt.Println(strconv.Itoa(count), " is length. So is ", len(letter))*/

	//Activity 21. Armstrong number
	/*for i := 1; i <= 500; i++ {
		x := i / 100
		y := (i % 100) / 10
		z := i % 10
		x = int(math.Pow(float64(x), 3))
		y = int(math.Pow(float64(y), 3))
		z = int(math.Pow(float64(z), 3))
		if i == (x + y + z) {
			fmt.Println(strconv.Itoa(i), " is an Armstrong number.")
		}

	}*/

}
