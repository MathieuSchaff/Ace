package main

import "fmt"

func main(){
	//declare a variable
	card := newCard()
	fmt.Println(card)
}

func newCard() string{
	return "Five of Diamonds"
}
 