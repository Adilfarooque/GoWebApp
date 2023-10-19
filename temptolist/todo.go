package temptolist

import (
	"html/template"
	"net/http"
)

type task struct {
	Name string
	Done bool
}

var tpl *template.Template

// To do list be exported
var Todo []task

func Todlist() {
	//Todo = []task{{"Wake Up Early", true}, {"Make ablution", true}, {"Pray Tahajjud", true}, {"Pick Groceries", false}, {"Make breakfast", false}, {"Iron dresses", false}, {"Workout at 9'O clock", false}}
	Todo = []task{{"Wake Up Early", true}, {"Make ablution", true}, {"Pray Tahajjud", true}, {"Pick Groceries", true}, {"Make breakfast", true}, {"Iron dresses", true}, {"Workout at 9'O clock", false}}
	tpl, _ = tpl.ParseGlob("temptolist/templates/*.html")
	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "list.html", Todo)
	})
	http.ListenAndServe(":30175", nil)
}
