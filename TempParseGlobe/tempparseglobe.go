package tempparseglobe

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func Glob() {
	tpl, _ = tpl.ParseGlob("TempParseGlobe/templates/*.html")
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", nil)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "login.html", nil)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "/about.html", nil)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "/contact.html", nil)
	})
}
