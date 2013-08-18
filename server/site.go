package main

import (
  "fmt"
  "log"
  "html/template"
  "io/ioutil"
  "net/http"
  "regexp"
  "os"
)

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
  p, err := loadPage(title)
  if err != nil {
    http.NotFound(w, r)
    return
  }
  renderTemplate(w, "view", p)
}

var templates = template.Must(template.ParseGlob("templates/*"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  err := templates.ExecuteTemplate(w, tmpl+".html", p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

const lenPath = len("/view/")

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]
    if !titleValidator.MatchString(title) {
      http.NotFound(w, r)
      return
    }
    fn(w, r, title)
  }
}

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(w, "hello, world!")
}

func main() {
  http.HandleFunc("/", hello)
  http.HandleFunc("/view/", makeHandler(viewHandler))
  // localhost: 5000
  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
