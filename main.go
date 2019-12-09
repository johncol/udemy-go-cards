package main

import "fmt"

func main() {
	deck := NewDeck()
	fmt.Println(" --- Deck")
	deck.print()

	deck, hand := Deal(deck, 5)
	fmt.Println(" --- Deck after deal")
	deck.print();
	fmt.Println(" --- Hand")
	hand.print();
}
