package main

import "fmt"

// Deck list of cards
type Deck []string

func (deck Deck) print() {
	for _, card := range deck {
		fmt.Println(card)
	}
}


// NewDeck builds a new Deck
func NewDeck() Deck {
	suits := []string{"Hearts","Spades","Dimonds","Clubs"}
	values := []string{"A","Two","Three","Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Q", "K"}

	deck := Deck{}

	for _, suit := range suits {
		for _, value := range values {
			card := value+" of "+suit
			deck = append(deck, card)
		}
	}

	return deck
}