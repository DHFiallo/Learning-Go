package main

import (
	"fmt"
	"time"
)

/*func goroutine(limit int) {
	for i := 0; i < limit; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("calling goroutine")
		fmt.Println(i)
	}
}

func anothergoroutine(limit int) {
	for i := 0; i < limit; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Println("calling anothergoroutine")
		fmt.Println(i)
	}
}*/

func message(ch chan string) {
	//Add a sleep delay to the channel
	time.Sleep(6 * time.Second)
	ch <- "Hello World" // we use ch followed by <- to write data to the channel
}

// calculate the max
func computeMax(ch chan int, numbers [4]int) {
	max := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	ch <- max
	//close(ch)
}

// calculate the min
func computeMin(ch chan int, numbers [4]int) {
	min := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	ch <- min
	close(ch)
}

func numberGenerator(ch chan int, limit int) {
	for i := 0; i < limit; i++ {
		ch <- i
	}
	close(ch)
}

func main() {

	ch := make(chan int)
	go numberGenerator(ch, 20)

	//What happens is that the method loops through 0 and 20, but the println
	//catches it inbetween and prints the number it's on at the time
	// read the first number
	b := <-ch
	fmt.Println("Number:", b)

	// read the second number
	b = <-ch
	fmt.Println("Number:", b)

	//This for loop is inbetween the loop in the method, so it prints out
	//while the method runs and changes b

	for b := range ch {
		fmt.Println("Number:", b)
	}

	/*numbers := [4]int{25, 64, 75, 5}
	fmt.Println(numbers)

	ch1 := make(chan int)
	go computeMax(ch1, numbers)
	b, ok := <-ch1
	fmt.Printf("Channel is closed: %v\n", ok)
	fmt.Printf("Max is: %v\n", b)

	ch2 := make(chan int)
	go computeMin(ch2, numbers)
	b, ok = <-ch2
	fmt.Printf("Channel is closed: %v\n", ok)
	fmt.Printf("Min is: %v\n", b)*/

	/*ch := make(chan string)
	fmt.Println(ch)
	go message(ch) //executes the goroutine with input as the channel
	b := <-ch
	fmt.Println(b)
	fmt.Println("This will execute last") //This waits due to channel imposing a block/wait
	*/

	/*fmt.Printf("Channel Type is %T", a)
	var a chan int
	var a = make(chan int)
	fmt.Println(a)
	go goroutine(10)
	go anothergoroutine(10)
	time.Sleep(6 * time.Second)
	fmt.Println("main goroutine")*/
}
