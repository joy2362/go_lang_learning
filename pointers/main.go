package main

import "fmt"

func main()  {
	welcome := "welcome"
	fmt.Println(welcome)


	// declear and default value of pointer 
	var ptr *int 
	fmt.Printf( "type of ptr is %v \n",ptr )


	//create a variable that point to a variable
	num := 123;

	var prt2 = &num  //& ref to some memory 

	fmt.Println(prt2) //return memory address
	fmt.Println(*prt2) // pointer value


	//action perform in actual value not copy of a value
	*prt2 = *prt2 +  2

	fmt.Println(num)
	
}