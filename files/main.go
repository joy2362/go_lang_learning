package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	welcome := "welcome";
	fmt.Println(welcome);

	content := "hello world";

	file , err := os.Create("./some.txt")

	checkNillError(err)

	length , err := io.WriteString(file, content)

	checkNillError(err)

	fmt.Println(length)

	defer file.Close()

	readFile("./some.txt")

}

//read file 
func readFile(fileName string){
	dataByte , err := os.ReadFile(fileName)

	checkNillError(err)

	fmt.Println(dataByte)
	fmt.Println(string(dataByte))
}

func checkNillError(err error){
	if err != nil{
		panic(err)
	}
}