package tempstruct

import (
	"html/template"
	"net/http"
)

type ProductSpec struct {
	Size        string
	Weight      float32
	Description string
}

type Product struct {
	ProID int
	Name  string
	Cost  float32
	Specs ProductSpec
}

var tpl *template.Template
var prod1 Product
func ProductInfo() {
	prod1 = Product{
		ProID: 22,
		Name: "I Phone 15 pro",
		Cost: 999,
		Specs: ProductSpec{
			Size: "15 x 7 x 7nm",
			Weight: 65,
			Description: "The Best Product ",
		},
	}

	tpl , _ = tpl.ParseGlob("/tempStruct/templates/*.html")
	http.HandleFunc("/prd1",prdHandlerFunc)
	http.ListenAndServe(":8080",nil)
}

func prdHandlerFunc(w http.ResponseWriter , r *http.Request)  {
	tpl.ExecuteTemplate(w,"prd1.html",prod1)
}