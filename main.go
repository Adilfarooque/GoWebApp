package main

import (
	"GoWebApp/testifelse"
	"fmt"
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
	testifelse.IfEles()
	fmt.Println("listen and serving on : 8000")
	err := http.ListenAndServe("localhost:8080", nil)
	fmt.Println("server error", err)
}
