package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	welcome := "welcome"
	fmt.Println(welcome)
	greeter()

	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001",r))
}

func greeter(){
	fmt.Println("hey there")
}

func serveHome(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("<h1>welcome</h1>"))
}