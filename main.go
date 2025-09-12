package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hello")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("I am a software engineering student")
}

func main(){
	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler) //route
	mux.HandleFunc("/about", aboutHandler) //route

	fmt.Println("server running on :3000")

	err := http.ListenAndServe(":3000", mux)  // Failed to start  the server
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}