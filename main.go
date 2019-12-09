package main

import "fmt"

func main() {
	deck := NewDeck()
	fmt.Println(" --- Deck")
	deck.Print()

	deck, hand := Deal(deck, 5)
	fmt.Println(" --- Deck after deal")
	deck.Print();
	fmt.Println(" --- Hand")
	hand.Print();
}
