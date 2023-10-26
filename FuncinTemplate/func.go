package funcintemplate

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var err error

var tpl, _ = template.New("index.html").Funcs(template.FuncMap{
	
	"CanCashpr": func(p float64) string {
		reminder := int(p*100) % 5
		quotiant := int(p*100) / 5

		if reminder < 3 {
			pr := float64(quotiant*5) / 100
			s := fmt.Sprintf("%.2f", pr)
			return s
		}
		pr := (float64(quotiant*5) + 5) / 100
		s := fmt.Sprintf("%2.f", pr)
		return s
	},
	"Upper": strings.ToUpper,
}).ParseFiles(".FuncinTemplate/templates/index.html")


var p float64

func UsingFunc() {

	fmt.Println("tpl.Tree", *tpl.Tree)
	p = 3.31
	tpl, _ = tpl.ParseFiles("index.html")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = tpl.ExecuteTemplate(w, "index.html", p)
		if err != nil {
			log.Fatal(err)
		}
	})
}
