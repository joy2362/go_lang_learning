package main

import "fmt"

func main()  {
	welcome := "welcome to struct"
	fmt.Println(welcome)

	//no inheritance in go lang , No super or parent class
	//capital name user for public variable

	joy := User{"joy" , "test@test.com" , true, 24}
	fmt.Println(joy)
	//to get details for struct we have to use %+v
	fmt.Printf("Joy details are: %+v \n" , joy)
	fmt.Printf("Joy name is: %v \n" , joy.Name)
	joy.GetStatus()
	joy.NewEmail()
}


type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Printf("Status is: %v \n" , u.Status)
}

func (u User) NewEmail(){
	u.Email = "test@test.com"
	fmt.Println(u.Email)
}