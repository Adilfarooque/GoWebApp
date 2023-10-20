package tempindex

import (
	"html/template"
	"net/http"
)

var tpl,_ = template.New("myTemplate").Funcs(template.FuncMap{
	"lastItem": func(s []string) string {
		lastIndex := len(s) - 1
		return s[lastIndex]
	},
}).ParseFiles("index.html")

var g []string

func IndexValue() {
	g = []string{"Milk", "Coffe Poweder", "Mint", "Pottato", "Cheese", "Flour", "Sugar", "Salt", "Brocooli"}
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", g)
	})
	http.ListenAndServe(":8080", nil)
}
