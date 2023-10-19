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

func LanguageTest() {
	//u = User{"Decaprio", "English", false}
	u = User{"Eman Gazi", "Arabic", true}
	// 	u = User{"Dojo Kat", "Mandarin", true}
	// 	u = User{"AK47", ">", true}
	tpl, _ = tpl.ParseGlob("tempifelse/templates/*.html")
	http.HandleFunc("/lang", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "membership.html", u)
	})
	http.ListenAndServe(":9090", nil)
}
