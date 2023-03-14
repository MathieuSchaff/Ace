package main

import "fmt"

func main() {
	//declare a variable
	// card := newCard()
	// fmt.Println(card)
	cards := deck{newCard(), "Ace of Spades"}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
