package main

import (
  "fmt"
  "log"
  "html/template"
  "io/ioutil"
  "net/http"
  //"regexp"
  "os"
)

type Page struct {
  Title string
  Body []byte
}


func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  p := &Page{Title: "index"}
  renderTemplate(w, "index", p)
}


var templates = template.Must(template.ParseGlob("templates/*"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  //idiomatic this
  err := templates.ExecuteTemplate(w, "site.html", p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  err = templates.ExecuteTemplate(w, tmpl+".html", p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

//const lenPath = len("/")

//var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

//func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
  //return func(w http.ResponseWriter, r *http.Request) {
    //title := r.URL.Path[lenPath:]
    //if !titleValidator.MatchString(title) {
      //http.NotFound(w, r)
      //return
    //}
    //fn(w, r, title)
  //}
//}

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(w, "hello, world!")
}

func main() {
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
  http.HandleFunc("/", indexHandler)
  //http.HandleFunc("/view/", makeHandler(viewHandler))
  // localhost: 5000
  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
