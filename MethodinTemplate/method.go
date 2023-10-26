package methodintemplate

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

// price of item
type Price float64

func (p Price) CanCashPr() string {
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
}

var p Price

func UsingMethod() {
	var err error
	p = 3.94
	tpl, err = tpl.ParseGlob("./FuncinTemplate/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.html", p)
		if err != nil {
			log.Fatal(err)
		}
	})
}
