package main

import "fmt"

func main() {
	printState()
	card := createNewCard()
	fmt.Println(card)

	cards := []string{"Five of Spade", createNewCard()}
	cards = append(cards, "Jack of Diamond")
	fmt.Println(cards)
	fmt.Println(len(cards))

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func createNewCard() string {
	return "Five of Diamonds"
}
