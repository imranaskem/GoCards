package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	data, e := ioutil.ReadFile(filename)

	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}

	return strings.Split(string(data), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for i := range d {
		newPosition := rng.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
