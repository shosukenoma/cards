package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	fmt.Println(cards)

	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}