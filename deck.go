package main

import "fmt"

// Deck list of cards
type Deck []string

type Hand = Deck

func (deck Deck) print() {
	for _, card := range deck {
		fmt.Println(card)
	}
}


// NewDeck builds a new Deck
func NewDeck() Deck {
	suits := []string{"Hearts","Spades","Dimonds","Clubs"}
	values := []string{"A", "J", "Q", "K"}

	deck := Deck{}

	for _, suit := range suits {
		for _, value := range values {
			card := value+" of "+suit
			deck = append(deck, card)
		}
	}

	return deck
}

// Deal takes N cards out of the deck
func Deal(deck Deck, cards int) (Deck, Hand) {
	return deck[cards:], deck[:cards]
}