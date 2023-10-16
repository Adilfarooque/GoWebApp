package main

import (
	"html/template"
	"net/http"
	
)

var tpl *template.Template

func main() {
	tpl, _ = template.ParseFiles("index1.html")
	// tpl, _ = template.ParseFiles("")
	// tpl, _ = template.ParseFiles("")
	// tpl, _ = template.ParseFiles("")
	http.HandleFunc("/", indexHandlerFunc)
	http.ListenAndServe("localhost:8080", nil)
}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

