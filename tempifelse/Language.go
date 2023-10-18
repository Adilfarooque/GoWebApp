package tempifelse

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

func Language() {
	//u = User{"Decaprio", "English", false}
	u = User{"Eman Gazi", "Arabic", true}
	// 	u = User{"Dojo Kat", "Mandarin", true}
	// 	u = User{"AK47", ">", true}
	tpl, _ = tpl.ParseGlob("tempifelse/templates/*.html")
	http.HandleFunc("/lang", langHandler)
	http.ListenAndServe(":8080", nil)
}

func langHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "lang.html", u)
}
