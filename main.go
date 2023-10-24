package main

import (
	tempparseglobe "GoWebApp/TempParseGlobe"
	"net/http"
)

func main() {
	//basicserver2.BasicServer2()
	// basicserver1.BaseServer()
	// tempindex.IndexValue()
	//templatesparsing.ParsingFiles()
	tempparseglobe.Glob()
	http.ListenAndServe(":8080", nil)
}
