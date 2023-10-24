package basicserver2

import (
	"fmt"
	"net/http"
)

func BasicServer2() {
	http.HandleFunc("/hello", helloHandlerFunc)
	http.HandleFunc("/about", aboutHandlerFunc)

}

func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.url.path: , %s", r.URL.Path)
}

func aboutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome :) GopherWorld!")
}
