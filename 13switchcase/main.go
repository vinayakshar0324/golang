package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in GO")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)
	fmt.Println("Value of Dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 nd you can open")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("dice value is 3")
	case 4:
		fmt.Println("dice value is 4")
	case 5:
		fmt.Println("dice value is 5")
	case 6:
		fmt.Println("You can roll dice again")
	default:
		fmt.Println("what was that!")
	}

}
