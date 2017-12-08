package main

import (
	"fmt"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"s", "h", "d", "c"}
	cardValues := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	hand := d[:handSize]

	return hand, d[handSize:]
}
