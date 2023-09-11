package main

import (
	"fmt"
	"sort"
)

func main()  {
	welcome := "welcome"
	fmt.Println(welcome)
	
	//take user input 
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Size of array:")
	// num , _ := reader.ReadString('\n')
	// arraySize , err := strconv.ParseInt(strings.TrimSpace(num) , 10 , 32)
	
	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Print(arraySize)
	// }


	//slice 
	var array = []string{}
	fmt.Printf("Type of array is %T\n" , array)

	var array2 = []string{"abc" ,"def" , "ghi"}
	fmt.Println(array2)

	array2 = append(array2, "test")
	fmt.Println( array2)
	array2 = append(array2[1:])
	fmt.Println( array2)

	array2 = append(array2[:2])
	fmt.Println( array2)

	array3 := make([]int , 4)

	array3[0] = 100
	array3[1] = 400
	array3[2] = 300
	array3[3] = 200

	array3 = append(array3, 500, 900, 700)
	fmt.Println( array3)

	sort.Ints(array3)

	fmt.Println( array3)
	fmt.Println( sort.IntsAreSorted(array3))

	//how to remove a value from slice 
	var lang = []string{"c" , "c++" ,"java" , "php", "javascript"}
	index := 2

	//clone index
	lang = append(lang[:index]  , lang[index+1:]...)
	fmt.Println(lang)

	
}