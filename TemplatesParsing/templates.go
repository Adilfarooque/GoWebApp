package templatesparsing

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func ParsingFiles() {
	//tpl, _ = template.ParseFiles("TemplatesParsing/index1.html")
	// tpl, _ = tpl.ParseFiles("TemplatesParsing/data1/index2.html")
	tpl, _ = tpl.ParseFiles("TemplatesParsing/data1/data2/index3.html")
	//tpl, _ = tpl.ParseFiles("..index4.html")
	http.HandleFunc("/", indexHandlerfunc)
}

func indexHandlerfunc(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
