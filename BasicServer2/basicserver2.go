package basicserver2

import (
	"fmt"
	"net/http"
)

func BasicServer2() {
	http.HandleFunc("/hello", helloHandlerFunc)
	http.HandleFunc("/wel", welcomeHandlerFunc)

}

func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello All!!!")
}

func welcomeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome :)")
}
