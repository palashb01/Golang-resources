package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - should be in a separate file

type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) Isempty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - learnCodeonlie.in")
	r := mux.NewRouter()
	// seeding of the data
	// have multiple courses
	courses = append(courses, []Course{{
		CourseID:    "2",
		CourseName:  "AngularJS Bootcamp",
		CoursePrice: 299,
		Author: &Author{
			FullName: "LearnCodeOnline.in",
			Website:  "LearnCodeOnline.in",
		},
	}, {
		CourseID:    "3",
		CourseName:  "VueJS Bootcamp",
		CoursePrice: 299,
		Author: &Author{
			FullName: "LearnCodeOnline.in",
			Website:  "LearnCodeOnline.in",
		},
	}}...)
	
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneController).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file
// serverHome route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my server</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)
	// loop through courses, find matching
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No courses found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.Isempty(){
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// generate a unique id, convert it into string
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses{
		if item.CourseID == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"] 
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}


// delete controller
func deleteOneController(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)

	for index, item := range courses{
		if item.CourseID == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
}