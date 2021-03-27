package main

import "fmt"

var externalCard string // variables can only be initialized outised a function

func main() {
	var card string = "Ace of Spades" // only string values can be assigned to this variable
	fmt.Println(card)
	/*
	Go is a "Static Types" language like C++, Java.
	Examples of "Dynamic Types" languages are: JavaScript, Ruby, Python

	Some basic Go types:
	- bool
	- string
	- int
	- float64

	It is possible to define variable types in different ways, also the examples below
	*/
	var card2 = "two of Spades"
	fmt.Println(card2)
	card3 := "three of Spades" // ':=' must be used only to initialize a variable
	fmt.Println(card3)
	card3 = "another three of Spades"
	fmt.Println(card3)
	externalCard = "external Ace of Spades"
	fmt.Println(externalCard)
}