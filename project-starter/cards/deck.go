package main

import "fmt"

type deck []string

//the convention for method in golang in here as below
//require the name of the receiver with the first letter of the type
//example deck : the first letter is d
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
