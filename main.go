package main

import (
	insertingifelse "GoWebApp/InsertingIfElse"
	"net/http"
)

func main() {
	//basicserver2.BasicServer2()
	// basicserver1.BaseServer()
	// tempindex.IndexValue()
	//templatesparsing.ParsingFiles()
	//tempparseglobe.Glob()
	//insertingdataintohtmltemplates.Insert()
	//structbase.InStructBase()
	insertingifelse.IfEles()
	http.ListenAndServe("localhost:8080", nil)
}
