package rangefunction3

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

var g []string

func Range3() {
	var err error
	g = []string{"Apple", "Orange", "Mango", "Banana", "Pinapple", "Shamam", "Burtukal"}
	tpl, err = tpl.ParseGlob("./RangeFunction3/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "groceries2.html", g)
		if err != nil {
			log.Fatal(err)
		}
	})

}
