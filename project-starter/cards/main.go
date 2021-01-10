package main

import "fmt"

func main() {
	printState()
	card := createNewCard()
	fmt.Println(card)
}

func createNewCard() string {
	return "Five of Diamonds"
}
