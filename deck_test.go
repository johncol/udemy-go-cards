package main

import (
	"testing"
	"strings"
)

func TestNewDeck_HasTheExpectedSize(t *testing.T) {
	deck := NewDeck()

	expectedDeckSize := 16
	actualDeckSize := len(deck)
	if actualDeckSize != expectedDeckSize {
		t.Errorf("Expected deck length is %v but got %v", expectedDeckSize, actualDeckSize)
	}
}

func TestNewDeck_Has4CardsForEachSuit(t *testing.T) {
	suits := []string{"Hearts","Spades","Dimonds","Clubs"}
	deck := NewDeck()
	expectedCardsPerSuit := 4

	for _, suit := range suits {
		suitCards := filterDeckBySuit(deck, suit)
		actualCardsPerSuit := len(suitCards)
		if actualCardsPerSuit != expectedCardsPerSuit {
			t.Errorf("Expected cards for suit %v was %v, but got %v", suit, expectedCardsPerSuit, actualCardsPerSuit)
		}
	}
}

func filterDeckBySuit(deck Deck, suit string) []string {
	suitCards := []string{}
	for _, card := range deck {
		if strings.HasSuffix(card, suit) {
			suitCards = append(suitCards, card)
		}
	}
	return suitCards
}