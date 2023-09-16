package main

import "fmt"

func main()  {
	welcome := "welcome"
	fmt.Println(welcome)
	days := []string{"Sunday" , "Saturday" , "Monday" , "Friday"}

	fmt.Println(days)

	for i:=0; i < len(days) ; i++{
		fmt.Println(days[i])
	}

	//foreach but i return index 
	for i := range days{
		fmt.Println(days[i])
	}
	//foreach with key and value 
	for i , day := range days{
		fmt.Printf("index is %v and value is %v \n" , i , day)
	}


	//while loop
	roughValue := 1

	for roughValue < 10{

		if roughValue ==5 {
			roughValue++
			continue
		}

		if roughValue == 2 {
			goto lco
		}

		fmt.Println(roughValue)
		roughValue++
	}

	lco: 
		fmt.Println("test go to syntex")
	
}