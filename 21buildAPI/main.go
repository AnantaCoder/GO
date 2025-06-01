package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// "math/rand"

	"github.com/gorilla/mux"
)

// how its look like =- model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"` //its has not  created so mark as a pointer 
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db
var courses []Course

//different methods  -- middlewares or helpers
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// func main() {

// }

//controllers - file 
//server home route 
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to the api </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses ")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One course ")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request 
	params := mux.Vars(r)
	//loop through the course and find id and return   the response 
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

//course controller add one course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate

	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append this  course into courses

	// You can call genetrateRandomID() here when needed

	course.CourseId = generateRandomID()
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func generateRandomID() string {
	b := make([]byte, 15)
	_, err := rand.Read(b)
	if err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return hex.EncodeToString(b)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")

	// Set the Content-Type header for the response.
	w.Header().Set("Content-Type", "application/json")
	// Grab ID from request parameters.
	params := mux.Vars(r)

	// Loop through the courses to find the one to update.
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Decode the request body into a new Course object.
			var updatedCourse Course
			err := json.NewDecoder(r.Body).Decode(&updatedCourse)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest) // Handle JSON decoding errors.
				return
			}

			// Preserve the original CourseId from the URL parameter.
			updatedCourse.CourseId = params["id"]

			// Replace the old course with the updated course in the slice.
			courses[index] = updatedCourse

			// Encode and send the updated course as a response.
			json.NewEncoder(w).Encode(updatedCourse)
			return // Exit the function after successful update.
		}
	}
	// If the loop completes without finding the course, send a "Course not found" response.
	json.NewEncoder(w).Encode(map[string]string{"message": "No course found with given id to update"})
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course ")
	w.Header().Set("Content-Type", "application/json")

	// grad id from request 
	params := mux.Vars(r)

	//loop through to find id of the course and delete it 
	for index, course := range courses {
		if course.CourseId == params["id"] {
			//delete course at index by appending slice bfr and after the index 
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Course with id " + params["id"] + " is deleted ",
			})
			return //exit string 
		}
	}

	// if course not found with given id 
	json.NewEncoder(w).Encode(map[string]string{
		"Message": "No course found with given id to delete ",
	})
}

func main() {
	r := mux.NewRouter()
	// routes 
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//courses apopend 
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "ReactJS",
		CoursePrice: "2999",
		Author: &Author{
			Fullname: "John Doe",
			Website:  "kaka.com",
		},
	})

	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "golang",
		CoursePrice: "4565",
		Author: &Author{
			Fullname: "Aditya Rikari",
			Website:  "johndoe.com",
		},
	})

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// go build .
