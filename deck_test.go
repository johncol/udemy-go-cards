package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	expectedDeckSize := 16
	actualDeckSize := len(deck)
	if actualDeckSize != expectedDeckSize {
		t.Errorf("Expected deck length is %v but got %v", expectedDeckSize, actualDeckSize)
	}
}