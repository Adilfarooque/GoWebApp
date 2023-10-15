package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl , _ = template.ParseFiles("index1.html")
	
	
	http.HandleFunc("/",indexHandlerFunc)

	http.ListenAndServe(":8080",nil)
}

func indexHandlerFunc(w http.ResponseWriter , r *http.Request){
	tpl.Execute(w,nil)
}