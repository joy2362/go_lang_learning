package main

import "fmt"

func main()  {
	welcome := "welcome"
	fmt.Println(welcome)

	//array 
	var array [4]string 

	array[0] = "apple"
	array[1] = "banana"
	array[3] = "string"

	// print array 
	fmt.Println("array list is ",array)

	// print array value one by one
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	
	// print array length
	fmt.Println("array list is ",len(array))

	// array 
	var array2 = [4]string{"test" , "test2" ,"test3" , "test4"} 

	fmt.Println("array list is ",array2)

	// print array value one by one
	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}
}