package main

import (
	"fmt"
	"io"
	"os"
)

func fileFn() {
	file, err := os.Create("./myFile.txt")
	if err != nil {
		panic(err)
	}

	lenOfStr, err := io.WriteString(file, "Here's my string take it")

	if err != nil {
		panic(err)
	}

	fmt.Printf("length of the string is %v", lenOfStr)

	defer file.Close() // defer keyword puts the statement at the end of the fn

	readFile("./myFile.txt")
}

func readFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}