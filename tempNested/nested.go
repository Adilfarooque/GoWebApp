package tempnested

import (
	"html/template"
	"net/http"
)

type task struct {
	Name string
	Done bool
}

var tpl *template.Template

func Nested() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", nil)
	})

	http.ListenAndServe(":8080", nil)
}
