package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/howeyc/fsnotify"
	"github.com/russross/blackfriday/v2"
)

// TODO: 404 page

type Page struct {
	Body template.HTML
}

func Markdown(s string) template.HTML {
	return template.HTML(blackfriday.Run([]byte(s)))
}

var (
	funcs = template.FuncMap{"md": Markdown}

	templates = template.Must(template.New("").Funcs(funcs).ParseGlob("templates/*.html"))

	// pre-render these for prod
	indexPage  = &Page{Body: renderContent("index")}
	resumePage = &Page{Body: renderContent("resume")} // TODO: remove md magic here
)

func renderTemplate(w http.ResponseWriter, data interface{}) {
	err := templates.ExecuteTemplate(w, "site.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, indexPage)
}

func handleResume(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, resumePage)
}

func handleContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := filepath.Base(r.URL.Path)
		if page == "/" {
			page = "index"
		}
		renderTemplate(w, &Page{Body: renderContent(page)})
	}
}

func renderContent(page string) template.HTML {
	file, _ := filepath.Glob("content/" + page + ".*")

	if len(file) > 0 {
		switch filepath.Ext(file[0]) {
		case ".md":
			return renderMarkdown(file[0])
		}
	}
	var dummy interface{}
	return renderHTML(page, &dummy)
}

// TODO return err
func renderHTML(page string, data interface{}) template.HTML {
	var buf bytes.Buffer
	templates.ExecuteTemplate(&buf, page+".html", data)
	return template.HTML(buf.Bytes())
}

// TODO return err
func renderMarkdown(page string) template.HTML {
	md, _ := ioutil.ReadFile(page)
	return template.HTML(blackfriday.Run(md))
}

// reload templates on modify
func listenForChanges() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				//recompile templates
				templates = template.Must(template.New("").Funcs(funcs).ParseGlob("templates/*.html"))
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("templates/")
	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}

func serveFile(url string, filename string) {
	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func main() {
	var port int
	var prod bool
	flag.IntVar(&port, "port", 5000, "port to run on")
	flag.BoolVar(&prod, "prod", false, "production")
	flag.Parse()

	serveFile("/favicon.ico", "./favicon.ico")
	serveFile("/sitemap.xml", "./sitemap.xml")
	serveFile("/robots.txt", "./robots.txt")

	if !prod {
		log.Println("running in debug mode")

		// handle dynamically
		http.HandleFunc("/", handleContent())

		go listenForChanges()

		// TODO return / listen and serve
	} else {

		log.Println("running in prod mode")

		// static functions, for memory/speed
		http.HandleFunc("/", handleIndex)
		http.HandleFunc("/hire/resume", handleResume)

		// TODO autocert

	}

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
