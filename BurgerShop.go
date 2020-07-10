//Darius Fiallo
//June 15th, 2020
//BurgerShop

package main

import (
	"fmt"
	"strconv"
)

type burger struct {
	bun     bool
	price   float64
	dressed bool
}

type drink struct {
	price float64
	name  string
}

type side struct {
	price float64
	name  string
}

type combo struct {
	b     burger
	s     side
	d     drink
	price float64 //Combined price of all 3 with a 20% discount

}

func user_input_burger() burger {
	b := new(burger)
	var user string
	fmt.Println("Do you want a bun with your burger? Y/N?")
	fmt.Scanln(&user)
	if user == "Y" {
		b.bun = true
		b.price += 5
	} else {
		b.bun = false
		b.price += 4
	}

	fmt.Println("Do you want any dressing? Y/N")
	fmt.Scanln(&user)
	if user == "Y" {
		b.dressed = true
	} else {
		b.dressed = false
	}

	return *b
}

func user_input_drink() drink {
	d := new(drink)
	var user string
	fmt.Println("Do you want a drink?\n1. Water\n2. Coke\n3. Pepsi\n4. Mountain Dew\nPut a number.")
	fmt.Scanln(&user)
	userInt, _ := strconv.Atoi(user)
	switch userInt {
	case 1:
		d.name = "water"
		d.price += 1
	case 2:
		d.name = "coke"
		d.price += 2
	case 3:
		d.name = "pepsi"
		d.price += 2
	case 4:
		d.name = "mountain dew"
		d.price += 2
	default:
		d.name = "none"
	}
	return *d
}

func user_input_side() side {
	s := new(side)
	var user string
	fmt.Println("Do you want a side?\n1. Fries\n2. Onion Rings\n3. Salad\n4. Cornbread Muffin\nPut a number.")
	fmt.Scanln(&user)
	userInt, _ := strconv.Atoi(user)
	switch userInt {
	case 1:
		s.name = "fries"
		s.price += 1
	case 2:
		s.name = "onion rings"
		s.price += 2
	case 3:
		s.name = "salad"
		s.price += 2
	case 4:
		s.name = "cornbread muffin"
		s.price += 2
	default:
		s.name = "none"
	}
	return *s
}

func user_input_combo() combo {
	c := new(combo)
	burg := user_input_burger()
	drnk := user_input_drink()
	sd := user_input_side()
	c.b = burg
	c.d = drnk
	c.s = sd
	c.price = (burg.price + drnk.price + sd.price) * 0.8

	return *c
}

func take_order_from_user() {
	var name string
	fmt.Println("What's the name for your order?")
	fmt.Scanln(&name)

	for {
		c := user_input_combo()
		var choice string
		fmt.Println("Another combo? Y/N")
		fmt.Scanln(&choice)
		if choice == "Y" {
			continue
		} else {
			fmt.Println(c)
			fmt.Println("Thank you for ordering at Darius's Burgers")
			break
		}
	}
}

func main() {
	take_order_from_user()

}
