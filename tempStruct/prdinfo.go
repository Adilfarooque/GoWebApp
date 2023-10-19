package tempstruct

import (
	"html/template"
	"net/http"
)

type productSpec struct {
	Size        string
	Weight      float32
	Description string
}

type product struct {
	ProID int
	Name  string
	Cost  float32
	Specs productSpec
}

var tpl *template.Template
var prod1 product

func ProductInfo() {
	prod1 = product{
		ProID: 22,
		Name:  "I Phone 15 pro",
		Cost:  999,
		Specs: productSpec{
			Size:        "15 x 7 x 7nm",
			Weight:      65,
			Description: "The Best Product ",
		},
	}

	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/productinfo", productInfoHandler)
	http.ListenAndServe(":8080", nil)
}

func productInfoHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "productinfo.html", prod1)
}
