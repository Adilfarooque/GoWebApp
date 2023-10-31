package uploadfile

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template
var err error

func UpLoad() {
	tpl, err = tpl.ParseGlob("./UploadFile/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		//if method is GET then load form, if not then upload successful msg
		if r.Method == "GET" {
			tpl.ExecuteTemplate(w, "fileUpload.html", nil)
			return
		}
		//func (r *Request) ParseMultipartFrom(maxMemory int64)error
		r.ParseMultipartForm(10)

		//func (r *Request) FormFile(key String)(multipart.File, *multipart.FileHeader, error)
		file, fileHeader, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Printf("FileHeader.Filename: %v\n", fileHeader.Filename)
		fmt.Printf("FileHeader.Size: %v\n", fileHeader.Size)
		fmt.Printf("FileHeader.Header: %v\n", fileHeader.Header)

		// tempFile , err :=ioutil.TempFile("image", "upload-*.png")

		contentType := fileHeader.Header["Content-Type"][0]
		fmt.Println("Content Type:", contentType)
		var osFile os.File
		//func TempFile(dir,pattern String)(f *os.file, error)
		
		
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the home page", nil)
	})

}
