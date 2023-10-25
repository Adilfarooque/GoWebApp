package testifelse

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var Tpl *template.Template

type User struct {
	Name   string
	Lang   string
	Member bool
}

var u User

func IfEles() {
	//u = User{"Rob Smith", "English", false}
	u = User{"Suarex", "Spanish", true}
	//u = User{"Elone","Mandarin",true}
	//u = User{"Bill007","?",true_
	fmt.Println("1")
	tpl, err := Tpl.ParseGlob("./testifelse/templates/*.html")
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	fmt.Println("2")
	Tpl = tpl
	http.HandleFunc("/welcome", welcomehandler)
}

func welcomehandler(w http.ResponseWriter, r *http.Request) {
	err := Tpl.ExecuteTemplate(w, "membership2.html", u)
	if err != nil {
		fmt.Println("Parsing error", err)
	}
}
