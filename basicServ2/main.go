package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandlerFunc)

	http.HandleFunc("/about", aboutHandleFunc)

	http.ListenAndServe(":8080", nil)
}

//Read and write and Read and Request
func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "r.URL.Path: , %S", r.URL.Path)
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome To Our Gopher World!")
}



