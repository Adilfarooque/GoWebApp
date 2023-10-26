package main

import (
	rangefunction "GoWebApp/RangeFunction"
	rangefunction2 "GoWebApp/RangeFunction2"
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
	//testifelse.IfEles()
	rangefunction.Range()
	rangefunction2.TodoList()
	fmt.Println("listen and serving on : 8000")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("server error", err)
}
