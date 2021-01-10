package main

import "fmt"

func main() {
	printState()
	card := createNewCard()
	fmt.Println(card)

	cards := deck{"Ace of Spades", createNewCard()}
	cards = append(cards, "Jack of Diamond")
	fmt.Println(cards)
	fmt.Println(len(cards))

	cards.print()
}

func createNewCard() string {
	return "Five of Diamonds"
}
