package main

import (
	"fmt"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"math/rand"
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

// SaveToFile saves the content of the deck as a string in a file
func (deck Deck) SaveToFile(filename string) error {
	bytes := []byte(deck.ToString())
	return ioutil.WriteFile(filename, bytes, 0666)
}

// Shuffle shuffles deck cards
func (deck Deck) Shuffle() {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(deck)

	for index := range deck {
		targetIndex := generator.Intn(length)
		deck[index], deck[targetIndex] = deck[targetIndex], deck[index]
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

// NewDeckFromFile builds a new deck using the contents of a file in the system
func NewDeckFromFile(filename string) Deck {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading deck from file, error was\n", err)
		os.Exit(1)
	}

	commaSeparatedCards := string(bytes)
	cards := strings.Split(commaSeparatedCards, ",")

	return Deck(cards)
}

// Deal takes N cards out of the deck
func Deal(deck Deck, cards int) (Deck, Hand) {
	return deck[cards:], deck[:cards]
}
