package main

import (
	basicserver2 "GoWebApp/BasicServer2"
	"net/http"
)

func main() {
	// basicserver1.BaseServer()
	basicserver2.BasicServer2()
	http.ListenAndServe(":8080", nil)
}
