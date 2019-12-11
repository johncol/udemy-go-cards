package main

import "fmt"

func main() {
	deck := NewDeck()
	fmt.Println(" --- Deck")
	deck.Print()

	fmt.Println(" --- Deck as string")
	fmt.Println(deck.ToString())

	filename := "deck.txt"
	fmt.Println(" --- Saving Deck to file", filename)
	deck.SaveToFile(filename)

	fmt.Println(" --- Reading deck from file", filename)
	deck = NewDeckFromFile(filename)
	deck.Print()

	deck, hand := Deal(deck, 5)
	fmt.Println(" --- Deck after deal")
	deck.Print();
	fmt.Println(" --- Hand")
	hand.Print();

	fmt.Println(" --- Deck after shuffle")
	deck.Shuffle()
	deck.Print();
}
