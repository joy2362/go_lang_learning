package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model for course - file
type Course struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Name string `json:"name"`
	Website string `json:"website"`
}

//fake db
var courses []Course

//middleware
func (c *Course) IsEmpty() bool {
	return c.Name == ""
}

func main() {
	welcome := "welcome"
	fmt.Println(welcome)

	r := mux.NewRouter()

	//seedeing
	courses = append(courses, Course{
		Id: "2",
		Name: "Go lang",
		Price: 5000,
		Author: &Author{
			Name: "Zahid",
			Website: "test123",
		},
	})

	courses = append(courses, Course{
		Id: "3",
		Name: "c++",
		Price: 1000,
		Author: &Author{
			Name: "Zahid123",
			Website: "test",
		},
	})
	

	//routing
	r.HandleFunc("/" , home).Methods("GET")
	r.HandleFunc("/course" , index).Methods("GET")
	r.HandleFunc("/course/{id}" , show).Methods("GET")
	r.HandleFunc("/course" , store).Methods("POST")
	r.HandleFunc("/course/{id}" , update).Methods("PUT")
	r.HandleFunc("/course/{id}" , destroy).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":8002", r))
}


//controller - file

//home
func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>welcome to go lang</h1>"))
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Println("get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
	return
}

func show(w http.ResponseWriter,r *http.Request){
	fmt.Println("get single courses")
	w.Header().Set("Content-Type","application/json")

	//grab id from request
	params :=  mux.Vars(r)

	//loop course match id and return response
	for _, course := range courses{
		if course.Id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return 	
		}
	}

	json.NewEncoder(w).Encode("No course found ")
	return 
}


func store(w http.ResponseWriter,r *http.Request){
	fmt.Println("save course")
	w.Header().Set("Content-Type","application/json")

	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	//what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return 
	}

	//TODO: Check only name is exits


	for _,item := range courses{
		if item.Name == course.Name {
			json.NewEncoder(w).Encode("Name already exits")
			return
		}
	}

	//generate unique id, string 
	//append course into course

	course.Id = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

}


func update(w http.ResponseWriter,r *http.Request){
	fmt.Println("update course")
	w.Header().Set("Content-Type","application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop,remvoe ,add with my id

	checkCourse := false

	for index,course := range courses{
		if course.Id == params["id"] {
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.Id = params["id"]

			courses = append(courses, course)

			checkCourse = true
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	if !checkCourse {
		json.NewEncoder(w).Encode("Course not found")
		return
	}
	 
}


func destroy(w http.ResponseWriter,r *http.Request){
	fmt.Println("Destroy course")
	w.Header().Set("Content-Type","application/json")

	//grab id from request
	params := mux.Vars(r)
	checkCourse := false
	for index,course := range courses{
		if course.Id == params["id"] {
			courses = append(courses[:index],courses[index+1:]...)
			checkCourse = true
			break
		}
	}
	if checkCourse {
		json.NewEncoder(w).Encode("Course removed")
		return
	}else{
		json.NewEncoder(w).Encode("course not found")
		return
	}
}