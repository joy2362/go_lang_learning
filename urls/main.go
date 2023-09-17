package main

import (
	"fmt"
	"net/url"
)

const MyUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	welcome := "welcome"
	fmt.Println(welcome)


	//parsing
	result , err := url.Parse(MyUrl)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	query := result.Query()

	for key, value := range query{
		fmt.Println(key , value)
	}

	partsOfUrl :=  &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "/user=joy",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
	
}