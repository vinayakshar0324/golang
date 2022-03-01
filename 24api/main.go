package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//model for courses
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db
var courses []Course

//middleware, helper file
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseID: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Vinayk Sharma", Website: "lco.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")



	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers - file

//serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request

	params := mux.Vars(r)

	//loop through courses find marching id and reutrn the response
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty(){
		json.NewEncoder(w).Encode("Please Send Some Data")
		return
	}

	//GENERATE THE UNIQUE ID, STRING
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return 

}

func updateOneCourse(w http.ResponseWriter, r *hhtp.Request)  {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")


	//first grab id from req
	param := mux.Vars(r)


	//loop, id, remove, add with my id 

	for index, course := range courses{
		if course.CourseID == param["id"]{
			courses = append(courses[:index], courses[index+1:]... )
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO: send a response when id is not found 

	func deleteOneCourse(w http.ResponseWriter, r *hhtp.Request)  {
		fmt.Println("Delete One Course")
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)


		//loop, id , remove (index, index+1)
		for index, course := range courses{
			if course.CourseID == params["id"]{
				courses = append(courses[:index], courses[index+1:]... )
				break
			}
		}





	}

}
}