//Darius Fiallo, Jasmin Smith, Andrew Horn
//June 19, 2020
//Cause a deadlock between dining philosophers

package main

import (
	"fmt"
	"time"
)

/*

    F5  P1
 P5        F1
F4         P2
  P4      F2
    F3 P3

*/

type philosopher struct {
	name      string
	LeftHand  bool
	RightHand bool
}

type fork struct {
	held bool
}

func (p *philosopher) tryToEat(right *fork, left *fork) {
	for {
		if right.held == false {
			right.held = true
			p.RightHand = right.held
		}
		if left.held == false {
			left.held = true
			p.LeftHand = left.held
		}
		if left.held && right.held {
			break
		}
		if p.LeftHand {
			fmt.Println("Left hand has a fork. Deadlock")
		} else if p.RightHand {
			fmt.Println("Right hand has a fork. Deadlock.")
		} else {
			fmt.Println("Back to thinking, loser. Try and get some forks.")
		}
	}
	if p.LeftHand && p.RightHand {
		go p.eat(left, right)
	}

}

func (p *philosopher) eat(left *fork, right *fork) {
	fmt.Println("Finally some good f***ing food. -" + p.name)
	time.Sleep(5 * time.Second)
	go p.setDownForks(left, right)
}

func (p *philosopher) setDownForks(left *fork, right *fork) {
	p.LeftHand = false
	left.held = false
	p.RightHand = false
	right.held = false
	fmt.Println("Forks from ", p.name, " have been set down.")
	go p.tryToEat(left, right)

}

func main() {

	/*

		    F5  P1
		 P5        F1
		F4         P2
		  P4      F2
		    F3 P3

	*/

	f1 := fork{false}
	f2 := fork{false}
	f3 := fork{false}
	f4 := fork{false}
	f5 := fork{false}

	p1 := philosopher{"Plato", false, false}
	p2 := philosopher{"Zhuangzi", false, false}
	p3 := philosopher{"Descartes", false, false}
	p4 := philosopher{"Kongzi", false, false}
	p5 := philosopher{"Nagarjuna", false, false}
	go p1.tryToEat(&f5, &f1)
	go p2.tryToEat(&f1, &f2)
	go p3.tryToEat(&f2, &f3)
	go p4.tryToEat(&f3, &f4)
	go p5.tryToEat(&f4, &f5)

	time.Sleep(60 * time.Second)
}

/*
//Darius Fiallo, Jasmin Smith, Andrew Horn
//June 19, 2020
//Cause a deadlock between dining philosophers

package main

import (
	"fmt"
	"time"
)

/*

    F5  P1
 P5        F1
F4         P2
  P4      F2
    F3 P3

*/
/*
func (p *philosopher) tryToEat(left *bool, right *bool) {
	for {
		for {
			if *left {
				fmt.Println(p.name, "picking up left fork")
				*left = false
				p.left = true
			}
			if *right {
				fmt.Println(p.name, "picking up right fork")
				*right = false
				p.right = true
			}
			if p.left && p.right {
				break
			}
		}
		fmt.Println(p.name, "is eating")
    time.Sleep(5 * time.Millisecond) // eating
		p.setDownForks(left, right)
	}
}


func (p *philosopher) setDownForks(left *bool, right *bool) {
	p.left = false
	p.right = false
	*left = true
	*right = true
	fmt.Println(p.name, "put down his forks")
}

func main() {
	f1 := true
	f2 := true
	f3 := true
	f4 := true
	f5 := true
	p1 := philosopher{"1Plato", false, false}
	p2 := philosopher{"2Zhuangzi", false, false}
	p3 := philosopher{"3Descartes", false, false}
	p4 := philosopher{"4Kongzi", false, false}
	p5 := philosopher{"5Nagarjuna", false, false}
	go p1.tryToEat(&f5, &f1)
	go p2.tryToEat(&f1, &f2)
	go p3.tryToEat(&f2, &f3)
	go p4.tryToEat(&f3, &f4)
	go p5.tryToEat(&f4, &f5)
	time.Sleep(1000 * time.Second)
}
*/
