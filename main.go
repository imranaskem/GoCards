package main

func main() {
	cards := newDeck()

	hand, remainingDeck := cards.deal(2)

	hand.print()
	remainingDeck.print()
}
