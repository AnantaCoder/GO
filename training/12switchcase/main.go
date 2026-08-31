package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case in golang")
	max := 6
	min := 1 
	number := rand.Intn(max-min+1)+min
	// fmt.Println("Random number is: ", number)

	switch number{
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}
}