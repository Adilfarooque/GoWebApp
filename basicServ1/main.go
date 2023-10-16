package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",helloHandlerFunc)
	http.ListenAndServe(":8080",nil)
}

func helloHandlerFunc(w http.ResponseWriter , r *http.Request){
	fmt.Fprint(w, "Hello, World")
}

