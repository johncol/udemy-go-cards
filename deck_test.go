package main

import (
	"testing"
	"strings"
	"os"
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

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "test-deck.txt"
	_ = os.Remove(testFileName)

	deck := NewDeck()
	err := deck.SaveToFile(testFileName)
	if err != nil {
		t.Fatal("SaveToFile returned error: ", err)
	}

	deckFromFile := NewDeckFromFile(testFileName)

	if len(deck) != len(deckFromFile) {
		t.Errorf("Original deck and deck from file are not the same")
	}
	
	_ = os.Remove(testFileName)
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