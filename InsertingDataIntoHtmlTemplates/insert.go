package insertingdataintohtmltemplates

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var name = "Pichai"

func Insert() {
	tpl, _ = tpl.ParseGlob("InsertingDataIntoHtmlTemplates/templates/*.html")
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "welcome.html", name)
	})
}
