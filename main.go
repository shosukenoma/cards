package main

// import "fmt"

func main() {
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// cards.print()

	// hand.print()
	// remainingDeck.print()

	cards.shuffle()
	cards.print()
}