package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func newDeck() Deck {
	cards := Deck{}
	cardSymbols := []string{"Spade", "Clover", "Heart", "Diamond"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, symbol := range cardSymbols {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+symbol)
		}
	}
	return cards
}

func (d Deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) Deck {
	byteString, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	return Deck(strings.Split(string(byteString), ","))
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
