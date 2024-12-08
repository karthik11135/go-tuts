package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var courses = [][]string{{"hey", "alrighty", "boooo"}, {"what is this", "ii dont't know bro"}}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", sendGreeting).Methods("GET")

	r.HandleFunc("/json", sendJsonData).Methods("GET")

	r.HandleFunc("/params/{id}", readParams).Methods("GET")

	r.HandleFunc("/post", postController).Methods("POST")

	r.HandleFunc("/query", getQueryStr).Methods("GET")

	http.ListenAndServe(":3000", r)
}

func sendGreeting(w http.ResponseWriter, req *http.Request) {
	//sending normal text
	w.Write([]byte("Hey boy how you"))
}	

func sendJsonData(w http.ResponseWriter, req *http.Request) {
	//sending headers
	w.Header().Set("Content-Type", "application/json")

	//sending json
	json.NewEncoder(w).Encode(courses)
}

func readParams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//getting params
	params := mux.Vars(r);

	id := params["id"]

	//converting string to integeer
	idNum, err := strconv.Atoi(id)
	fmt.Println(id)

	if err != nil {
		json.NewEncoder(w).Encode("Send a number macha")
		return
	}

	if idNum < 2 {
    	json.NewEncoder(w).Encode(courses[idNum])
		return
	}
	json.NewEncoder(w).Encode("Something went wrong")
}


//getting data as a post requuest, a bit counterintuitive
type Person struct { 
    Name string `json:"name"` 
    Age int `json:"age"`
}

func postController(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("No body is sent")
		return
	}
	var person Person;
	json.NewDecoder(r.Body).Decode(&person);

	fmt.Println(person)
}

func getQueryStr(w http.ResponseWriter, r *http.Request) {
	hasFirst := r.URL.Query().Has("first")
	fmt.Println(hasFirst)
	first := r.URL.Query().Get("first")
	fmt.Println(first)
}