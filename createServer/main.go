package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	welcome := "welcome"
	fmt.Println(welcome)
	PerformGetRequest()
	PerformPostRequest()
	PerformPostFornRequest()
}

func PerformGetRequest(){
	const URL = "http://localhost:1111/get"

	res , err :=http.Get(URL)
	
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println("status code: " , res.StatusCode)
	fmt.Println("Content length: " , res.ContentLength)

	// dataByte , err :=io.ReadAll(res.Body)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(dataByte))

	var responseString strings.Builder

	contentByte , err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	contentCount , err := responseString.Write(contentByte)
	if err != nil {
		panic(err)
	}
	fmt.Println(contentCount)
	fmt.Println(responseString.String())
}

func PerformPostRequest(){
	const URL = "http://localhost:1111/post"

	//payload
	requestBody := strings.NewReader(`
		{
			"name" : "joy",
			"age" : "1234"
		}

	`)

	res , err := http.Post(URL , "application/json" , requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content , err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func PerformPostFornRequest(){
	const URL = "http://localhost:1111/post"

	//payload
	data := url.Values{}
	data.Add("name","joy")
	data.Add("age","123")
	data.Add("email","test@test.com")

	res , err := http.PostForm(URL , data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content , err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}