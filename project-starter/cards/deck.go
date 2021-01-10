package main

import "fmt"

type deck []string

func createNewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//the convention for method in golang in here as below
//require the name of the receiver with the first letter of the type
//example deck : the first letter is d
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
