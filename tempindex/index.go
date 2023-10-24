package tempindex

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	var err error

	// Create a template function
	tpl, err = template.New("myTemplate").Funcs(template.FuncMap{
		"lastItem": func(s []string) string {
			lastIndex := len(s) - 1
			return s[lastIndex]
		},
	}).ParseFiles("tempindex/templates/index.html")

	// log error
	if err != nil {
		log.Println(err)
	}
}

var g []string

func IndexValue() {
	g = []string{"Milk", "Coffe Poweder", "Mint", "Pottato", "Cheese", "Flour", "Sugar", "Salt", "Brocooli"}

	tpl, _ = tpl.ParseGlob("tempindex/templates/*.html")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", g)
	})
}
