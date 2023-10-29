package basicserver3

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
var err error

func BServer3() {
	tpl, err = tpl.ParseGlob("./BasicServer3/templates/*.html")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		err = tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	//myDir := http.Dir("C:/Users/ADIL FAR Q/Desktop/GoWebApp/BasicServer3/fileServer/Public")

	//func FilerServer(root FileSystem) Handler
	/*Handler--> myHandler := http.FileServer(myDir)
	//http.Handle("/", myHandler)
	*/

	/*Using Absolute Path
	http.Handle("/", http.FileServer(http.Dir("C:/Users/ADIL FAR Q/Desktop/GoWebApp/BasicServer3/fileServer/Public")))
	*/

	/*Using Relative Path
	  It Usually Support but it's did't work for me :(
	http.Handle("/", http.FileServer(http.Dir("./Public")))
	*/

	/*Does not work , will look at ./public/public
	http.Handle("/Public", http.FileServer(http.Dir("./Public")))
	*/

	/*Using http.StripPrefix to alter request before FileServer sees it.
	http.Handle("/Public/", http.StripPrefix("/Public/", http.FileServer(http.Dir("./Public"))))
	*/
	//http.Handle("/Public", http.FileServer(http.Dir(".")))
	

}
