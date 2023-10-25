package insertingifelse

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

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
	//u = User{"Bill007","?",true}

	tpl, _ = tpl.ParseGlob("InsertingIfElse/templates/*.html")
	http.HandleFunc("/welcome", welcomehandler)
}

func welcomehandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "membership.html", u)
}
