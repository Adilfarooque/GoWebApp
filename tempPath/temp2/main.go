package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	// tpl, _ = template.ParseFiles("index1.html")
	// tpl, _ = template.ParseFiles("data1/index2.html")
	tpl, _ = tpl.ParseFiles("../index4.html")
	// tpl, _ = template.ParseFiles("../index4.html")

	http.HandleFunc("/", indexHandlerFunc)
	http.ListenAndServe(":8080", nil)
}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
