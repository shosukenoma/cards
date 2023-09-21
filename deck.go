package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
// This allows us to create custom functions for the deck
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// use _ for variables that we won't be using later.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// d in this case is kinda like "this" or "self"
// it refers to the current working copy of the type deck.
// By convention, receiver arguments are defined using 1 or 2 letter var names.
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}