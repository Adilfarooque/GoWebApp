package tempindex

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

var g []string

func IndexValue() {
	g = []string{"Milk", "Coffe Poweder", "Mint", "Pottato", "Cheese", "Flour", "Sugar", "Salt", "Brocooli"}
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", g)
	})
	http.ListenAndServe(":8080", nil)
}
