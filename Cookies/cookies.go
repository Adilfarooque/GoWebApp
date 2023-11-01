package cookies

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool 			// not accessible by JavaScript , only sent to server
	SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
*/

var tpl *template.Template

func CookiesWorld() {
	tpl, err := tpl.ParseGlob("./Cookies/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//looks for cookie on our machine
		//func (r *request) Cookie(name string)(*Cookie, error)
		cookie, err := r.Cookie("2st-cookie")
		fmt.Println("cookie:", cookie, "err:", err)
		if err != nil {
			fmt.Println("cookie was not found")
			cookie = &http.Cookie{
				Name:     "2st-cookie",
				Value:    "my second cookie value",
				HttpOnly: true,
			}
			//func SetCookie(w ResponseWriter , cookies *cookie)
			http.SetCookie(w, cookie)
		}
		tpl.ExecuteTemplate(w, "index.html", nil)
	})

}
