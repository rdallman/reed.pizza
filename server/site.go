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
  html, _:= ioutil.ReadFile("templates/index.html")
  t := template.HTML(html)
  p := &Page{Title: "index", Body: t}
  renderTemplate(w, "index", p)
}

func hireHandler(w http.ResponseWriter, r *http.Request) {
  html, _:= ioutil.ReadFile("templates/hire.html")
  t := template.HTML(html)
  p := &Page{Title: "hire", Body: t}
  renderTemplate(w, "hire", p)
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
  md, _:= ioutil.ReadFile("assets/resume.md")
  t := template.HTML(blackfriday.MarkdownCommon(md))
  p := &Page{Title: "resume", Body: t}
  renderTemplate(w, "resume", p)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
  html, _:= ioutil.ReadFile("templates/projects.html")
  t := template.HTML(html)
  p := &Page{Title: "projects", Body: t}
  renderTemplate(w, "projects", p)
}

func shoeHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("templates/shoes.html")
  t.ExecuteTemplate(w, "shoes.html", &Page{})
  //html, _:= ioutil.ReadFile("templates/shoes.html")
  //t := template.HTML(html)
  //p := &Page{Title: "shoes", Body: t}
  //renderTemplate(w, "shoes", p)
}

//var templates = template.Must(template.ParseGlob("templates/*.html"))
var templates = template.Must(template.ParseFiles("templates/site.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  //idiomatic this
  err := templates.ExecuteTemplate(w, "site.html", p)
  //err := templates.ExecuteTemplate(w, tmpl+".html", p)
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

func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(w, "hello, world!")
}

func main() {
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/projects", projectsHandler)
  //serveSingle("/projects", "./templates/projects.html")
  http.HandleFunc("/hire", hireHandler)
  http.HandleFunc("/hire/resume", resumeHandler)
  // localhost: 5000
  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
