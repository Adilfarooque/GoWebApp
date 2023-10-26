package indexesingolang

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	// var err error
	tpl, _ = template.New("myTemplate").Funcs(template.FuncMap{
		"lastItem": func(s []string) string {
			lastIndex := len(s) - 1
			return s[lastIndex]
		},
	}).ParseFiles("./indexesingolang/templates/index.html")

	// if err != nil{
	// 	log.Fatal(err)
	// }
}

var g []string

func Indexesintemplates() {
	var err error
	g = []string{"milk", "tomato", "potato", "honey", "green beans", "flour", "sugar", "broccoli"}
	tpl, err = tpl.ParseGlob("./indexesingolang/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", g)
	})
}
