package main

import (
	// "crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	// "time"

	"github.com/gorilla/mux"
)

//creating the structure for the courses
type Course struct {
	CourseId string `json:"course_id"`
	CourseName string `json: "coursename"`
	CoursePrice string `json: "price"`
	Author *Author `json: "author"`
}
type Author struct {
	Authorid string `json:"author_id"`
	FullName string `json: "author_name"`
	AuthorEmail string `json: "author_email"`
	AuthorPhone string `json: "author_phone"`
}
var database []Course;

//checking if the courses are empty or not
// func (c *Course) isEmpty() bool{
// 	return c.CoursId == 0 && c.CourseName == "" && c.CoursePrice == ""
// }

func main(){
	fmt.Println("Backend Starts From Here");
	r := mux.NewRouter();
	
	//Seeders
	database = append(database, Course{CourseId: "1", CourseName: "Sidharth Learn", CoursePrice: "100", Author: &Author{FullName: "Sidharth", Authorid: "12", AuthorEmail: "choudharysidharth082000@gmail.com", AuthorPhone: "8448605993"}});
	database = append(database, Course{CourseId: "2", CourseName: "Anushka Learn", CoursePrice: "100", Author: &Author{FullName: "Anushka", Authorid: "11", AuthorEmail: "choudharysidharth082000@gmail.com", AuthorPhone: "8448605993"}});

	//routing 
	r.HandleFunc("/", serveHome).Methods("GET");
	r.HandleFunc("/getAllCourses", getAllCourses).Methods("GET");
	r.HandleFunc("/getAllCourses/{id}", getOneCourse).Methods("GET");
	r.HandleFunc("/deleteCourse/{id", deleteCourse).Methods("DELETE");
	r.HandleFunc("/postCourse", createCourse).Methods("POST");


	//listening to the port 
	log.Fatal(http.ListenAndServe(":4000", r));
	
}

//controllers
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to the home page</h1>"));
}
//getting all the courses
func getAllCourses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database)
}
//getting the sinfle course
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Getting one course")
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);
	for _,course := range database{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return;
		}
	}
	fmt.Println("Unable to find the course");
	json.NewEncoder(w).Encode("No Course Found in the Database");
}
// //posting the courses in the database
func createCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("This is the line for posting the data in the database");
	w.Header().Set("Content-Type", "application/json");
	//check if the request is empty or not
	if r.Body == nil {
		json.NewEncoder(w).Encode("No File Entered");
	}
	//checking the body is empty or not
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course);
	//things to be performed in this controller
	course.CourseId = strconv.Itoa(100);
	//append in the course
	database = append(database, course);

	//sending the response bakc
	json.NewEncoder(w).Encode("File Added Success");
	return;
}

//deleting the course
func deleteCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	//getting the id from the request
	params := mux.Vars(r);

	//looping through and delete the element
	for index,element := range database {
		if element.CourseId == params["id"]{
			//deleting the course
			database = append(database[:index], database[index+1:]...);
			break;
		}
	}
	//sending the response
	json.NewEncoder(w).Encode("File Deletion Sucecss");
}