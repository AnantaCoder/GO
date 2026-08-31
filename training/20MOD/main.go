package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome so mod in go")
	greater()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET") // HandleFunc is used to register a handler function for a specific route and HTTP method.

	// http.ListenAndServe(":8080", r) -> we can do this classic method 
	log.Fatal(http.ListenAndServe(":8000",r)) //throws  error  
}

func greater() {
	fmt.Println("Greater: Hey go users")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my server</h1>"))
}
