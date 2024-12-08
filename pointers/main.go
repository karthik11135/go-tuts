package main

import "fmt"

// 0x13323232a
// 0x13323232a <- pointer stores this
// 0x13323232a
// 0x13323232a
// 0x13323232a
// 0x13323232a
// 0x13323232a

// *ptr gives the actual value
// &var gives the memory adress

func changeNum(x *string) {
	*x = "Sathvi"
	return
}

func main() {
	var ptr *int; // pointer responsible to hold addresses of an integer variable
	fmt.Println(ptr) // nil because it doesnt store anything
	myNum := 23
	ptr = &myNum
	// 0x01 - 23 (myNum)
	// 0x02 - 0x01 (ptr)

	// *ptr gives 23
	// &myNum gives 0x01
	// ptr gives 0x01

	name := "Karthik"
	fmt.Println(name)
	changeNum(&name)
	fmt.Println(name)
}