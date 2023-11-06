package main

import (
	cookies "GoWebApp/Cookies"
	"context"
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
	// rangefunction.Range()
	// rangefunction2.TodoList()
	//rangefunction3.Range3()
	//indexesingolang.Indexesintemplates()
	// nestedtemplate.Nested()
	// methodintemplate.UsingMethod()

	//funcintemplate.UsingFunc();;;;;;;;

	//form.FormValid()
	//form2.FormValid2()

	//----Serving Files----
	//basicserver3.BServer3()

	//Uploading files to Golang server
	//uploadfile.UpLoad()

	//HTTP Cookies
	cookies.CookiesWorld()

	fmt.Println("listen and serving on : 8080")
	err := http.ListenAndServe(":8080", context.ClearHandle(http.DefaultServeMux))
	fmt.Println("server error", err)
	
}
