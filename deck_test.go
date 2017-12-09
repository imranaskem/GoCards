package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(cards))
	}

	if cards[0] != "2s" {
		t.Errorf("Expected first card is 2s but got %v", cards[0])
	}

	if cards[len(cards)-1] != "Ac" {
		t.Errorf("Expected last card is Ac but got %v", cards[len(cards)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := readFromFile(fileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 got %v", len(loadedDeck))
	}

	os.Remove(fileName)
}
