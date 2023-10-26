package rangefunction2

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type task struct {
	Name string
	Done bool
}

var tpl *template.Template

var Todo []task

func TodoList() {
	var err error
	Todo = []task{{"Wakeup early", true}, {"Pray Fajr", true}, {"Recite Quran", true}, {"Do Ariticulation", true}, {"Do Workout Every Morning", false}, {"Make breakfast", false}}
	tpl, err = tpl.ParseGlob("./RangeFunction2/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "todo.html", Todo)
		if err != nil {
			fmt.Println("Parsing error", err)
		}
	})
}
