package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"username"`
	Age int
	College string
	Cgpa float32 
}

func createJsonFn() {

	students := []User{
		{"Karthik", 23, "NITT", 8.63},
		{"Naveen", 24, "IITR", 9.55},
	}

	// Convert to JSON
	jsonData, err := json.Marshal(students)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Print JSON as a string
	fmt.Println(string(jsonData))
}