package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func (d deck) shuffle() {
	// Preperation for generating a truly random number
	// ref: https://pkg.go.dev/math/rand#New

	// We will use time UnixNano to get a different int64 value everytime
	// ref: https://pkg.go.dev/time#Time.UnixNano
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Swapping values
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}