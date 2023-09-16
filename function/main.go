package main

import "fmt"

func main()  {
	welcome := "welcome"
	fmt.Println(welcome)
	welcomeMessage()
	fmt.Println(add(1,3))
	ans , message := sub(1,3,5)
	fmt.Println(ans)
	fmt.Println(message)
}

func welcomeMessage()  {
	fmt.Println("welcome")
}

func add(a int , b int) int{
	return a + b
}

func sub(values ...int)(int , string){
	total := 0
	for _,val := range values {
		total -= val
	}

	return total , "done"
}