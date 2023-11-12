package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	file := http.FileServer(http.Dir("templates"))
	mux.Handle("/", file)
	http.ListenAndServe(":8080", mux)
}
