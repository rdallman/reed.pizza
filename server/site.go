package main

import (
  "fmt"
  "log"
  "html/template"
  "io/ioutil"
  "net/http"
  //"regexp"
  "github.com/russross/blackfriday"
  "os"
)

type Page struct {
  Title string
  Body template.HTML
}



//func loadPage(title string) (*Page, error) {
  //filename := title + ".txt"
  //body, err := ioutil.ReadFile(filename)
  //if err != nil {
    //return nil, err
  //}
  //return &Page{Title: title, Body: body}, nil
//}


func indexHandler(w http.ResponseWriter, r *http.Request) {
  p := &Page{Title: "index"}
  renderTemplate(w, "index", p)
}

func hireHandler(w http.ResponseWriter, r *http.Request) {
  p := &Page{Title: "hire"}
  renderTemplate(w, "hire", p)
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
  md, _:= ioutil.ReadFile("assets/resume.md")
  t := template.HTML(blackfriday.MarkdownCommon(md))
  p := &Page{Title: "resume", Body: t}
  renderTemplate(w, "resume", p)
}


var templates = template.Must(template.ParseGlob("templates/*"))
//var templates = template.Must(template.ParseFiles("templates/index.html", "templates/site.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  //idiomatic this
  templates.ExecuteTemplate(w, "site.html", p)
  err := templates.ExecuteTemplate(w, tmpl+".html", p)
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
  http.HandleFunc("/hire", hireHandler)
  http.HandleFunc("/hire/resume", resumeHandler)
  //http.HandleFunc("/view/", makeHandler(viewHandler))
  // localhost: 5000
  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
