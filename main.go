package main

import (
	insertingdataintohtmltemplates "GoWebApp/InsertingDataIntoHtmlTemplates"
	"net/http"
)

func main() {
	//basicserver2.BasicServer2()
	// basicserver1.BaseServer()
	// tempindex.IndexValue()
	//templatesparsing.ParsingFiles()
	//tempparseglobe.Glob()
	insertingdataintohtmltemplates.Insert()
	http.ListenAndServe(":8080", nil)
}
