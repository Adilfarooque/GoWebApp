package main

import (
	"fmt"
	"net/http"
)

type customHandler struct {
	m string
}

func (c customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(c.m)
}

func main() {
	mux:=http.NewServeMux()
	c1:= &customHandler{"Amjad"}
	c2:= &customHandler{"Adil"}
	c3:= &customHandler{"Ashfak"}
	

	mux.Handle("/Amjad",c1)
	mux.Handle("/Adil",c2)
	mux.Handle("/Ashfak",c3)

	http.ListenAndServe(":8080",mux)
}
