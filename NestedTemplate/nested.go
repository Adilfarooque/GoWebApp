package nestedtemplate

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

// type task struct{
// 	Name string
// 	Done bool
// }

func Nested() {
	var err error
	tpl, err = tpl.ParseGlob("./NestedTemplate/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "index2.html", nil)
		if err != nil {
			log.Fatal(err)
		}
	})

}
