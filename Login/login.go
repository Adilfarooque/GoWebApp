package login

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template
var db *sql.DB

func LoginUp() {
	var err error
	tpl, err = tpl.ParseGlob("./Login/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("*****registerAuthHandlerRunning*****")
		r.ParseForm()
		username := r.FormValue("username")
		//check username for only alphaNumeric
		var nameAlphaNumeric = true
		for _, char := range username {
			//func IsLetter(r rune)bool,func IsNumber(r rune)bool
			//if !unicode.IsLetter(char)&& unicode.IsNumber
			if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
				nameAlphaNumeric = false
			}
		}

		//check username pswdLength
		var nameLength bool
		if 5 <= len(username) && len(username) >= 50 {
			nameLength = true
		}

		//check password criteria
		password := r.FormValue("password")
		fmt.Println("password:", password, "\npswdLength:", len(password))
		//variable that must pass for password creation criteria
		var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdSpaces bool
		pswdSpaces = true

		for _, char := range password {
			switch {
			// func IsLower(r rune)bool
			case unicode.IsLower(char):
				pswdLowercase = true
			// func IsUpper(r rune)bool
			case unicode.IsUpper(char):
				pswdUppercase = true
				// func IsNumber(r rune)bool
			case unicode.IsNumber(char):
				pswdNumber = true
				// func IsPunct(r rune)bool , // func IsSymbol(r rune)bool
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				pswdSpecial = true
				//func IsSpace (r rune)bool,type rune = int32
			case unicode.IsSpace(int32(char)):
				pswdSpaces = false

			}
		}

		if 11 < len(password) && len(password) < 60 {
			pswdLength = true
		}

		if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdSpaces || !pswdLength {
			tpl.ExecuteTemplate(w, "register.html", "Please check username and password criteria")
			return
		}

		//check if username already exists for availability
		stmt := "SELECT UserID FROM bcrypt WHERE username = ?"
		row := db.QueryRow(stmt, username)

		var uID string
		err := row.Scan(&uID)

		if err != sql.ErrNoRows {
			fmt.Println("username already exists , err:", err)
			tpl.ExecuteTemplate(w, "register.html", "username already taken")
			return
		}

		//create hash from password
		var hash []byte

		//
		hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			//fmt.Println("bcrypt err:",err)
			tpl.ExecuteTemplate(w, "register.html", "There was a problem registering account")
			return
		}
		fmt.Println("hash:", hash)
		fmt.Println("String(hash):", string(hash))

		//func (db *DB)prepare(query sting)(*stmt, error)
		var insertStmt *sql.Stmt
		insertStmt, err = db.Prepare("INSERT INTO bcrypt(username , Hash) VALUES (?,?);")
		if err != nil {
			fmt.Println("Error Prepareing Statement :", err)
			tpl.ExecuteTemplate(w, "register.html", "Ther was a problem registering account")
			return
		}
		defer insertStmt.Close()

		var result sql.Result
		//func (s *stmt)Exces(arg ...interface{})(result,error)
		result, err = insertStmt.Exec(username, hash)
		rowsaff, _ := result.RowsAffected()
		lastIns, _ := result.LastInsertId()
		fmt.Println("rowsAff:", rowsaff)
		fmt.Println("lastIns:", lastIns)
		fmt.Println("err", err)
		if err != nil {
			fmt.Println("error inserting new user")
			tpl.ExecuteTemplate(w, "reginster.html", "There was a problem registering accoutn")
			return
		}
		fmt.Fprint(w, "congrats , your account has been successfully created")
	})

}

// loginHandler serves
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("username:", username, "password:", password)
	//retrieve password from db to compare (hash) wiht user supplied password's hash
	var hash string
	stmt := "SELECT Hash FROM bcrypt WHERE Username = ? "
	row := db.QueryRow(stmt, username)
	err := row.Scan(&hash)
	fmt.Println("hash from db:", hash)
	if err != nil {
		tpl.ExecuteTemplate(w, "login.html", "Check username and password")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	//return nill on success
	if err == nil {
		fmt.Fprint(w, "You have successfully logged in :)")
	}

	fmt.Println("incorrect password")
	tpl.ExecuteTemplate(w, "login.html", "check username and password")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****reginsterHandler running*****")
	tpl.ExecuteTemplate(w, "register.html", nil)
}
