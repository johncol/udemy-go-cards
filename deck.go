package main

import "fmt"

// Deck list of cards
type Deck []string

func (deck Deck) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}
