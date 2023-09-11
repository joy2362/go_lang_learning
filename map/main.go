package main

import "fmt"

func main()  {
	welcome := "welcome to map "
	fmt.Println(welcome);
	

	//create map with make synctex
	lang := make(map[string]string)

	lang["PHP"] = "php"
	lang["JAVA"] = "java"
	lang["JAVASCRIPT"] = "javascript"
	lang["RUBY"] = "RUBY"
	lang["C"] = "c"
	lang["C++"] = "c++"


	//print all
	fmt.Println(lang)

	//print each item 
	fmt.Println(lang["PHP"])
	fmt.Println(lang["JAVA"])
	fmt.Println(lang["JAVASCRIPT"])
	fmt.Println(lang["RUBY"])
	fmt.Println(lang["C"])
	fmt.Println(lang["C++"])


	// delete value from map
	delete(lang ,"PHP")
	fmt.Println(lang)

	//loops 

	//comma ok syntex
	// :=
	for key , value := range lang{
		fmt.Println(key)
		fmt.Println(value)

	}

		
}