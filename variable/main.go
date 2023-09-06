package main

import (
	"fmt"
)

const LogedInUser bool = true

func main()  {
	//string 
	var username string = "joy2362"
	fmt.Println(username)
	fmt.Printf("variable type: %T \n" , username)

	//boolean
	var isUser bool = true
	fmt.Println(isUser)
	fmt.Printf("variable type: %T \n" , isUser)

	//integer
	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("variable type: %T \n" , smallVal)

	//float
	var smallFloatVal float64 = 255.31123
	fmt.Println(smallFloatVal)
	fmt.Printf("variable type: %T \n" , smallFloatVal)

	//default value and some aliases
	var anotherVar int;
	fmt.Println(anotherVar)
	fmt.Printf("variable type: %T \n" , anotherVar)

	//implicit type
	var web = "test"
	fmt.Println(web)

	//no var style volaras operator outside function not allow
	numberOfUser := 3900
	fmt.Println(numberOfUser)
	fmt.Printf("variable type: %T \n" , numberOfUser)

	
	fmt.Println(LogedInUser)
	fmt.Printf("variable type: %T \n" , LogedInUser)

}