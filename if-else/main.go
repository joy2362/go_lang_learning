package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	welcome := "welcome"
	fmt.Println(welcome)
	//condition
	loginCount := 7

	if loginCount < 5 {
		fmt.Println("Too many attamped!!")
	} else if loginCount < 10 && loginCount > 5{
		fmt.Println("Login successfully. but careful next!!")
	} else {
		fmt.Println("Login successfully!!")
	}

	//switch case
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you can open")
	case 2 : 
		fmt.Println("you can move 2 spot")
	case 3 : 
		fmt.Println("you can move 3 spot")
		fallthrough
	case 4 : 
		fmt.Println("you can move 4 spot")
		fallthrough // it print next case with out condition check
	case 5 : 
		fmt.Println("you can move 5 spot")
	case 6 : 
		fmt.Println("you can move 6 spot")
	default: 
		fmt.Println("What was that")

	}


}