package basicserver1

import (
	"fmt"
	"net/http"
)

func BaseServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Pichai ! Its Me Aadi")
	})
}
