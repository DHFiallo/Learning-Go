//Darius Fiallo
//June 23, 2020
//Grocery Store

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func shop(queue chan int, wg *sync.WaitGroup) {

	for {
		cust, ok := <-queue
		if !ok {
			break
		}
		seed := rand.NewSource(time.Now().UnixNano())
		gen := rand.New(seed)

		fmt.Println("Customer ", cust, " is shopping.")
		wait := gen.Intn(56) + 5
		time.Sleep(time.Duration(wait) * time.Second)
		fmt.Println("Customer ", cust, " is done.")
		wg.Done()
	}

}

func generateShopper(buffer chan int, wg *sync.WaitGroup) {

	queue := make(chan int, 5)

	for {
		cust, ok := <-buffer
		if !ok {
			break
		}
		queue <- buffer
	}
	/*for i := 1; i <= 5; i++ {
		wg.Add(1)
		queue <- i
	}
	for i := 5; i <= num; i++ {
		wg.Add(1)
		queue <- i
		wait := gen.Intn(6) + 5
		time.Sleep(time.Duration(wait) * time.Second)
	}*/

}

func main() {

	//Make a buffer queue for this to go into
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)

	num := gen.Intn(51) + 50

	for i := 1; i <= num; i++ {
		buffer <- i
	}

	var wg sync.WaitGroup

	go shop(queue, &wg)
	go shop(queue, &wg)
	go shop(queue, &wg)
	go shop(queue, &wg)
	go shop(queue, &wg)

	generateShopper(buffer, &wg)

	wg.Wait()
	close(queue)
}

/*
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// With covid 19, our small grocery store can only allow
// 5 people in it at once, but we're very popular.
// People shop for between 5 and 60 seconds, and a new
// person shows up every 5-10 seconds. Between 50 and 100
// people show up every day. Use a WaitGroup, and go routines
// to simulate this grocery store printing out as people enter
// and leave the store, and enter the queue

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Store opens.")
	store := make(chan struct{}, 50)
	// allow first 5 people in
	for i := 0; i < 5; i++ {
		store <- struct{}{}
	}
	people := rand.Intn(51) + 50
	for i := 0; i < people; i++ {
		wg.Add(1)
		go person(i, store, &wg)
		time.Sleep(time.Duration(rand.Intn(6)+5) * time.Millisecond)
	}

	wg.Wait()
	fmt.Println("All people have arrived and left.")
}

func person(i int, store chan struct{}, wg *sync.WaitGroup) {
	var noWait bool
	select {
	case <-store:
		noWait = true
		personEnterStore(i)
		store <- struct{}{}
	default:
		fmt.Println("Person", i, "is waiting in the queue.")
	}
	if !noWait {
		<-store
		personEnterStore(i)
		store <- struct{}{}
	}
	wg.Done()
}

func personEnterStore(i int) {
	fmt.Println("Person", i, "is entering the store.")
	time.Sleep(time.Duration(rand.Intn(56)+5) * time.Millisecond)
	fmt.Println("Person", i, "is leaving the store.")
}
*/
