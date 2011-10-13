package hello

import (
    "template"
    "http"
    "fmt"
    "crypto/md5"
    "browserid"
)

type Page struct {
  Title string
  User string
  Gravatar string
  Login string
}

var user string = "Anonymous"

func init() {
    http.HandleFunc("/", handlerMain)
    http.HandleFunc("/verify", handlerVerify)
    http.HandleFunc("/logout", handlerLogout)
}

func fetch_gravatar(user string) string {
  var h = md5.New()
  h.Write([]byte(user))
  return fmt.Sprintf("%x", h.Sum())  
}

func render_page(w http.ResponseWriter) {  
  login := "<img src='i/sign_in_blue.png'>"
  display_user := "Anonymous"
  if user != "Anonymous" {
    login = ""
    display_user = user + " | <a href='/logout'>Logout</a>"
  }
  
  page := &Page { Title: "BrowserID in GO", 
                  User: display_user, 
                  Gravatar: fetch_gravatar(user), 
                  Login: login } 
  
  t, _ := template.ParseFile("pages/index.html", nil)
  err := t.Execute(w, page)
  
  if err != nil {
    http.Error(w, err.String(), http.StatusInternalServerError)  
  }

}

func handlerMain(w http.ResponseWriter, r *http.Request) {
  render_page(w)
}

func handlerVerify(w http.ResponseWriter, r *http.Request) {
  user = browserid.AppEngineVerify(r)
  fmt.Fprint(w, user)
}

func handlerLogout(w http.ResponseWriter, r *http.Request) {
  user = "Anonymous"
	http.Redirect(w, r, "/", http.StatusFound)
}

