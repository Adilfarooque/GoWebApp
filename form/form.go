package form

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type User struct {
	Username string
	Data     string
}

func FormValid() {
	var err error
	tpl, err = tpl.ParseGlob("./form/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/getform", func(w http.ResponseWriter, r *http.Request) {
		err = tpl.ExecuteTemplate(w, "getform.html", nil)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.HandleFunc("/processget", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("processGetHandler running")
		var s User
		s.Username = r.FormValue("Username")
		s.Data = r.FormValue("d	ata")
		fmt.Println("Username:", s.Username, "Data:", s.Data)
		err = tpl.ExecuteTemplate(w, "thanks.html", s)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		err = tpl.ExecuteTemplate(w, "postform.html", nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/processpost", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("processGetHandler running")
		var s User
		s.Username = r.FormValue("Username")
		s.Data = r.FormValue("d	ata")
		fmt.Println("Username:", s.Username, "Data:", s.Data)
		err = tpl.ExecuteTemplate(w, "thanks.html", s)
		if err != nil {
			log.Fatal(err)
		}
	})
}
