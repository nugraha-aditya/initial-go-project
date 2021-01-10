package main

import "fmt"

func main() {
	cards := createNewDeck()
	hand, remainingHand := cards.deal(3)

	fmt.Println("In hand cards: ")
	hand.print()

	fmt.Println("--------------------------")

	fmt.Println("Remaining hand cards: ")
	remainingHand.print()
}
