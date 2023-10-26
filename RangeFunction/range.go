package rangefunction

import (
	"html/template"
	"log"
	"net/http"
)

type GroceriesList []string

var tpl *template.Template

var g GroceriesList

func Range() {
	var err error
	g = GroceriesList{"Date", "Banana", "Ladies Finger", "Carrot", "Pinapple", "Apple", "Grapes", "Tomatto", "Chillie", "Orange"}
	tpl, err = tpl.ParseGlob("./RangeFunction/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "groceries.html", g)
	})
}
