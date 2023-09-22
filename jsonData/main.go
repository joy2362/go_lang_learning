package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	welcome := "welcome"
	fmt.Println(welcome)
	encodeJson()

	decodeJson()

}

type user struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	Password string `json:"-"`
	Hobby []string `json:"hobby,omitempty"`
	//Hobby []string `json:"hobby"`
}

func encodeJson(){
	users := []user{
		{"zahid" , "test@test.com" , "male" , "abc123" , []string{"code","dev"}},
		{"abdullah" , "abdullah@test.com" , "male" , "abc123" , []string{"code","dev"}},
		{"joy" , "joy@test.com" , "male" , "abc123" , nil},
	}

	//package this data as json
	output , err := json.MarshalIndent(users , "" , "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",output)
}

func decodeJson(){
	jsonDataFromUrl := []byte(`
		{
			"name": "zahid",
			"email": "test@test.com",
			"gender": "male",
			"hobby": ["code","dev"]
		}
	`)

	var userData user 	

	checkData := json.Valid(jsonDataFromUrl)

	if checkData {
		fmt.Println("json is valid")

		json.Unmarshal(jsonDataFromUrl , &userData)

		fmt.Printf("%#v" , userData)
	}else{
		fmt.Println("Json not valid")
	}

	fmt.Printf("\n")

	//add some extra key value
	var userDataGet map[string]interface{}

	json.Unmarshal(jsonDataFromUrl , &userDataGet)


	//fmt.Printf("%#v\n" , userDataGet)
	for k,v  := range userDataGet {
		fmt.Println(k, v)
	}

	


}