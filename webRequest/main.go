package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://cameronjournal.com/?feed=rss2"

func main() {
	welcome := "welcome"
	fmt.Println(welcome)
 
	res , err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response type is ", res.Status)
	fmt.Printf("%T \n" , res)

	defer res.Body.Close()

	dataByte , err := io.ReadAll(res.Body)
	
	if err != nil {
		panic(err)
	}

	content := string(dataByte)

	fmt.Println(content	)

}