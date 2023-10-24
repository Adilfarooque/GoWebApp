package main

import (
	templatesparsing "GoWebApp/TemplatesParsing"
	"net/http"
)

func main() {
	//basicserver2.BasicServer2()
	// basicserver1.BaseServer()
	// tempindex.IndexValue()
	templatesparsing.ParsingFiles()
	http.ListenAndServe(":8080", nil)
}
