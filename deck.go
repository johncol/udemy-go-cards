package main

import "fmt"

// Deck list of cards
type Deck []string

func (deck Deck) print() {
	for _, card := range deck {
		fmt.Println(card)
	}
}
