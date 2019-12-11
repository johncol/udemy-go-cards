package main

import (
	"fmt"
	"strings"
)

// Deck list of cards
type Deck []string

// Hand alias for Deck
type Hand = Deck

// Print prints all cards in a deck
func (deck Deck) Print() {
	for _, card := range deck {
		fmt.Println(card)
	}
}

// ToString builds string with all cards separated by commas
func (deck Deck) ToString() string {
	cards := []string(deck)
	return strings.Join(cards, ",")
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
