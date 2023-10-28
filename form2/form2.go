package form2

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Sub struct {
	Username string
	Num      int
	Myfloat  float64
	Updates  bool
}

var tpl *template.Template

var err error

func FormValid2() {
	tpl, err = tpl.ParseGlob("./form2/templates/*.html")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/getform", func(w http.ResponseWriter, r *http.Request) {
		err = tpl.ExecuteTemplate(w, "getform.html", nil)
		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("/processget", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ProcessgetHandler running..")
		var s Sub
		s.Username = r.FormValue("username")
		//Can't use r.FormValue("numberName")
		num := r.FormValue("numberName")
		//converted the num value into string
		//ASCII to int
		//func Atoi(s string) (int , errors)
		s.Num, _ = strconv.Atoi(num)
		s.Num = s.Num * 2

		//func ParseFloat(s string, bitSize int)
		s.Myfloat, err = strconv.ParseFloat(r.FormValue("FloatName"), 64)
		if err != nil {
			log.Fatal("error parsing float64")
		}
		if r.FormValue("upName") == "true" {
			s.Updates = true
		} else if r.FormValue("upName") == "false" {
			s.Updates = false
		}

		tpl.ExecuteTemplate(w, "thanks.html", s)
	})


	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		err = tpl.ExecuteTemplate(w, "postform.html", nil)
	})

	http.HandleFunc("/processpost", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ProcessPostHandler running..")
		var s Sub
		s.Username = r.FormValue("username")
		//Can't use r.FormValue("numberName")
		num := r.FormValue("numberName")
		//converted the num value into string
		//ASCII to int
		//func Atoi(s string) (int , errors)
		s.Num, _ = strconv.Atoi(num)
		s.Num = s.Num * 2

		//func ParseFloat(s string, bitSize int)
		s.Myfloat, err = strconv.ParseFloat(r.FormValue("FloatName"), 64)
		if err != nil {
			log.Fatal("error parsing float64")
		}
		if r.FormValue("upName") == "true" {
			s.Updates = true
		} else if r.FormValue("upName") == "false" {
			s.Updates = false
		}

		tpl.ExecuteTemplate(w, "thanks.html", s)
	})
}
