package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)
	switch diceNumber{
	case 1:
		fmt.Println("Dice number is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
		fallthrough
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
	}
}