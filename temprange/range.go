package temprange

import (
	"html/template"
	"net/http"
)

type GroceryList []string

var tpl *template.Template

var g GroceryList

func GroceriesList() {
	g = GroceryList{"Mlik", "Eggs", "Banana", "Cheese", "Flour", "Sugar", "Brocooli"}
	tpl, _ = template.ParseGlob("temprange/templates/*.html")
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "groceries.html", g)
	})
	http.ListenAndServe(":8080", nil)
}
