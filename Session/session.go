package session

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var tpl *template.Template

// func NewCookieStore(keyPairs ...[]byte) *CookieStore
var store = sessions.NewCookieStore([]byte("super-secret-password"))

func Session() {
	var err error
	tpl, err = tpl.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/create",createHandler)
	http.HandleFunc("/delete",deleteHandler)
	http.HandleFunc("/about",aboutHandler)
	
}
