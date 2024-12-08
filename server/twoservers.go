package main

import "net/http"

func twoServers() {
	mux := http.NewServeMux()

	serverOne := &http.Server{
		Addr: ":3000", 
		Handler: mux,
	}

	serverTwo := &http.Server{
		Addr: ":3001", 
		Handler: mux,
	}

	serverOne.ListenAndServe()
	serverTwo.ListenAndServe()
}