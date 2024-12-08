package main

import "fmt"

// struct
type Building struct {
	cost int
	place string
	kind string
}

//method
func (b Building) demolish() {
	fmt.Println("Building is demolisehed")
}

func structs() {
	// initialisation
	var school Building = Building{240000, "New Jersey", "School"}
	school.demolish()
}