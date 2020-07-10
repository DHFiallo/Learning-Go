//Darius Fiallo
//June 21, 2020
//Internet Cafe

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//For some reason, an array that's twice the size is needed.
var computers = [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}

func checkComputers(m *sync.Mutex) bool {
	m.Lock()
	for i := 0; i < len(computers); i++ {
		if computers[i] == false {
			computers[i] = true
			m.Unlock()
			return true
			break
		}
	}
	m.Unlock()
	return false
}

func getOffComputer(m *sync.Mutex) bool {
	m.Lock()
	for i := 0; i < len(computers); i++ {
		if computers[i] == true {
			computers[i] = false
			m.Unlock()
			return true
			break
		}
	}
	m.Unlock()
	return false
}

func tourist(name int, m *sync.Mutex) {
	if checkComputers(m) == false {
		fmt.Println("Tourist ", name, " waiting for turn.")
	}
	for {
		if checkComputers(m) {
			fmt.Println("Tourist ", name, " is online.")
			seed := rand.NewSource(time.Now().UnixNano())
			gen := rand.New(seed)

			num := gen.Intn(11) + 1
			time.Sleep(time.Duration(num) * time.Second)
			fmt.Println("Tourist", name, " is done, having spent "+strconv.Itoa(num)+" minutes online.")
			getOffComputer(m)
			break
		}
	}
}

/* This is my tourist function that when with the channels
func tourist(ch <-chan int, results chan<- string, name string) {
	_, ok := <-ch
	time.Sleep(time.Millisecond)
	if ok == false {
		fmt.Println("Tourist ", name, " waiting for turn.")
	}
	for {
		if ok {
			fmt.Println("Tourist ", name, " is online.")
			seed := rand.NewSource(time.Now().UnixNano())
			gen := rand.New(seed)

			num := gen.Intn(11) + 1
			time.Sleep(time.Duration(num) * time.Second)
			var res string = "Tourist" + name + " is done, having spent " + strconv.Itoa(num) + " minutes online."
			results <- res
		}
	}

}

func bufferTourists(ch chan int) {

}*/

func main() {

	/* This was where I tried to use channels
	computers := make(chan int, 8)
	results := make(chan string, 8)

	for i := 1; i <= 25; i++ {
		go tourist(computers, results, strconv.Itoa(i))
	}

	for i := 1; i <= 8; i++ {
		computers <- i
	}

	for i := 1; i <= 25; i++ {
		result := <-results
		fmt.Println(result)
	}
	time.Sleep(time.Second * 60)*/

	var m sync.Mutex
	fmt.Println(computers)
	for i := 1; i <= 25; i++ {
		go tourist(i, &m)
	}
	time.Sleep(time.Second * 60)
}
