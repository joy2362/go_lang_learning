package main

import (
	"fmt"
	"time"
)

func main()  {
	welcome := "welcome "
	fmt.Println(welcome)


	//current time
	//format required 02-01-2006 15:04:05 Monday
	presentTime := time.Now()

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	//custom time

	createdDate := time.Date(2021 , time.October , 10 ,10 , 10 , 0 , 0 , time.Local)

	fmt.Println(createdDate.Format("02-01-2006 15:04:05 Monday"))
}