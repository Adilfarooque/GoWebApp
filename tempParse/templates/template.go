	package templates

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func SectionTemplates() {
	tpl, _ = template.ParseGlob("tempParse/templates/*.html")
	http.HandleFunc("/", indexHandlerFunc)
	http.HandleFunc("/about", aboutHandlerFunc)
	http.HandleFunc("/contact", contactHandlerFunc)
	http.HandleFunc("/login", loginHandlerFunc)
	http.ListenAndServe(":8080", nil)
}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func aboutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)
}

func contactHandlerFunc(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "contact.html", nil)
}

func loginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}
