package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)



func sendReqFn() {
	res, err := http.Get("https://hisdasd.free.beeceptor.com")
	if err != nil {
		fmt.Println(("something ewnt wrong"))
		panic(err)
	}

	defer res.Body.Close()

	// byte data is retrieved
	data, _ := io.ReadAll(res.Body)


	//1. using string builder way
	var stringBuilder strings.Builder
	stringBuilder.Write(data)

	fmt.Println(stringBuilder.String())

	// 2. converting it into a string
	fmt.Println(string(data))
}