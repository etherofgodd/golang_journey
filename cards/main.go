package main

import "fmt"

func main() {
	cards := deck{"3 kings", returnCardValue()}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func returnCardValue() string {
	card := "Five of Diamonds"

	return card
}
