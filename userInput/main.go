package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "welcome To user input "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name:")

	// comma ok || error ok 
	name , _ := reader.ReadString('\n')

	fmt.Print("Your name is:",name)


}