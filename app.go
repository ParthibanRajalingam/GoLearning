package main

import "fmt"

func main() {
	firstVar := "Hello Go!"
	fmt.Println(firstVar)
	//Slice
	cards := newDeck()
	//Slice - Append
	//loop
	cards.print()
	cards.shuffle()
	fmt.Println(cards.toString())
	//cards.saveToFile("555Cards.txt")

	newFileDeck().print()
}

func newFileDeck() Deck {
	newDeck := newDeckFromFile("555Cards.txt")
	return newDeck
}
