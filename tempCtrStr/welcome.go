package tempctrstr

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var name = "Aadil"

func AddingToHtml() {
	tpl,_ = tpl.ParseGlob("tempCtrStr/templates/*.html")
	http.HandleFunc("/welcome",welcomeHandlerFunc)

	http.ListenAndServe(":8080",nil)
}

func welcomeHandlerFunc(w http.ResponseWriter , r *http.Request){
	tpl.ExecuteTemplate(w,"welcome.html",name)
}