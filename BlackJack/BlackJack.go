//Darius, Syed, Rahul
//June 17, 2020
//BlackJack

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type card struct {
	value  string
	number float64
	suit   string
}
type cardCollection interface {
	findCard(card *card) (bool, int)
	sendCard(card *card, target cardCollection)
	receiveCard(card *card)
}
type deck struct {
	cards []*card
	index int
}

func GenerateDeck() *deck {
	arr := make([]*card, 52)
	index := 0
	for i := 0; i <= 3; i++ {
		var suit string
		switch i {
		case 0:
			suit = "spades"
		case 1:
			suit = "hearts"
		case 2:
			suit = "clubs"
		case 3:
			suit = "diamonds"
		}
		for j := 1; j <= 13; j++ {
			a := card{}
			if j == 1 {
				a.value = "ace"
			} else if j == 13 {
				a.value = "king"
			} else if j == 12 {
				a.value = "queen"
			} else if j == 11 {
				a.value = "jack"
			} else {
				a.value = strconv.Itoa(j)
			}
			a.suit = suit
			a.number = float64(j)
			arr[index] = &a
			index++
		}
	}
	return &deck{arr, 0}
}
func (d *deck) checkComplete() bool {
	m := make(map[string][]card)
	for _, value := range d.cards {
		m[value.suit] = append(m[value.suit], *value)
	}
	return len(m["spades"]) == 13 && len(m["hearts"]) == 13 && len(m["diamonds"]) == 13 && len(m["clubs"]) == 13
}
func (d *deck) shuffleDeck() {
	if true {
		len1 := 0
		cards := make([]*card, 52)
		for len1 < 52 {
			var card *card
			var index int
			for card == nil {
				ns := rand.NewSource(time.Now().UnixNano())
				generator := rand.New(ns)
				index = generator.Intn(52)
				card = d.cards[index]
			}
			d.cards[index] = nil
			cards[len1] = card
			len1++
		}
		d.cards = cards
	}
}
func (d *deck) findCard(card *card) (bool, int) {
	for index, value := range d.cards {
		if value == card {
			return true, index
		}
	}
	return false, 53
}
func (d *deck) sendTopCard(collection cardCollection) {
	d.sendCard(d.cards[d.index], collection)
	d.index++
}
func (d *deck) sendCard(card *card, collection cardCollection) {
	found, index := d.findCard(card)
	if found {
		d.cards[index] = nil
		collection.receiveCard(card)
	}
}
func (d *deck) receiveCard(card *card) {
	d.cards = append(d.cards, card)
}
func (d *deck) distributeDeck(human2 *human, computer2 *computer) {
	d.sendTopCard(human2)
	d.sendTopCard(computer2)
	d.sendTopCard(human2)
	d.sendTopCard(computer2)
}

type player interface {
	turn(deck *deck)
	sum() float64
}
type human struct {
	cards []*card
}

func newHuman() *human {
	var cards []*card
	return &human{cards}
}
func (d *human) turn(deck *deck) {
	turns := 0
	for d.sum() <= 21 && turns != 2 {
		fmt.Printf("Sum is: %f\n", d.sum())
		str := userInput("Choose to Hit or Stay: ")
		if str == "hit" {
			deck.sendTopCard(d)
		} else if str == "stay" {
			break
		}
		turns++
	}
}
func (d *human) sum() float64 {
	sum := 0.0
	for _, value := range d.cards {
		sum += value.number
	}
	return sum
}
func (d *human) findCard(card *card) (bool, int) {
	for index, value := range d.cards {
		if value == card {
			return true, index
		}
	}
	return false, 53
}
func (d *human) sendCard(card *card, collection cardCollection) {
	found, index := d.findCard(card)
	if found {
		d.cards[index] = nil
		collection.receiveCard(card)
	}
}
func (d *human) receiveCard(card *card) {
	d.cards = append(d.cards, card)
}

type computer struct {
	cards []*card
}

func newComputer() *computer {
	var cards []*card
	return &computer{cards}
}

func (d *computer) findCard(card *card) (bool, int) {
	for index, value := range d.cards {
		if value == card {
			return true, index
		}
	}
	return false, 53
}
func (d *computer) sendCard(card *card, collection cardCollection) {
	found, index := d.findCard(card)
	if found {
		d.cards[index] = nil
		collection.receiveCard(card)
	}
}
func (d *computer) receiveCard(card *card) {
	d.cards = append(d.cards, card)
}
func (d *computer) turn(deck *deck) {
	for d.sum() < 17 {
		deck.sendTopCard(d)
	}
}
func (d *computer) sum() float64 {
	sum := 0.0
	for _, value := range d.cards {
		sum += value.number
	}
	return sum
}
func userInput(input string) string {
	var response string
	fmt.Print(input)
	fmt.Scanln(&response)
	return response
}
func main() {
	d := GenerateDeck()
	d.shuffleDeck()
	c := newComputer()
	h := newHuman()
	d.distributeDeck(h, c)
	c.turn(d)
	h.turn(d)
	if h.sum() > 21 && c.sum() <= 21 || (c.sum() == 21 && h.sum() != 21) || c.sum() > h.sum() {
		fmt.Println("YOU LOSE")
	} else if c.sum() > 21 && h.sum() <= 21 || (h.sum() == 21 && c.sum() != 21) || h.sum() > c.sum() {
		fmt.Println("YOU WIN")
	} else {
		fmt.Println("TIE")
	}
}
