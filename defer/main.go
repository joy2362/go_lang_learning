package main

import "fmt"

func main()  {
	welcome := "welcome"
	fmt.Println(welcome)


	//defer last in first out 
	defer fmt.Println("world")
	defer fmt.Println("world2")
	defer fmt.Println("world3")
	fmt.Println("hello")
	myDefer()


}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}