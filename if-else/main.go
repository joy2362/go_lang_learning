package main

import "fmt"

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
}