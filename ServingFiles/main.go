package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

var err error

func main() {
	tpl, err = tpl.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	//create our new var myDir at type http.Dir
	myDir := http.Dir("./D:/C DRIVE/Desktop/GoWebApp/ServingFiles/public")
	//func FileServer(root FileSystem)Handler
	myHandler := http.FileServer(myDir)
	http.Handle("/", myHandler)

	//Using absolute path
	//http.Handle("/", http.FileServer(http.Dir("./C:/Users/USER/OneDrive/Desktop/Example/ServingFiles/public")))

	//Using relative path
	//http.Handle("/", http.FileServer(http.Dir("./public")))

	//does not work, will look at ./public/public
	//http.Handle("/public", http.FileServer(http.Dir("./public")))

	//use http.StripPrefix to alter request before FileServer sees it
	// http.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "hello.html", nil)
	})

	http.ListenAndServe(":8080", nil)
}
