package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	welcome := "welcome to convert user input type"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name:")

	name , _ := reader.ReadString('\n')
	fmt.Printf("Your name is: %v \n" ,name)

	fmt.Print("Enter your age:")
	age , _ := reader.ReadString('\n')
	fmt.Printf("Your age is: %v \n" ,age)

	ageConverted , err := strconv.ParseInt(strings.TrimSpace(age) , 10 , 32)


	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(ageConverted)
	}


	fmt.Printf("Your age is: %v \n" ,ageConverted)
	
}