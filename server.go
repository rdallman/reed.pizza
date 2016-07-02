package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/howeyc/fsnotify"
	"github.com/russross/blackfriday"
)

type Page struct {
	Title string
	Body  template.HTML
}

func Markdown(s string) template.HTML {
	return template.HTML(blackfriday.MarkdownCommon([]byte(s)))
}

var funcs = template.FuncMap{"md": Markdown}

var templates = template.Must(template.New("").Funcs(funcs).ParseGlob("templates/*.html"))

func renderTemplate(w http.ResponseWriter, data interface{}) {
	err := templates.ExecuteTemplate(w, "site.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := filepath.Base(r.URL.Path)
		if page == "/" {
			page = "index"
		}
		renderTemplate(w, &Page{Title: page, Body: renderContent(page)})
	}
}

func renderContent(page string) template.HTML {
	file, _ := filepath.Glob("content/" + page + ".*")

	//TODO idiomatic?
	if len(file) > 0 {
		switch filepath.Ext(file[0]) {
		case ".md":
			return renderMarkdown(file[0])
		case ".json":
			return renderJSON(file[0])
		}
	}
	return renderHTML(page, new(interface{}))
}

//TODO return err
func renderHTML(page string, data interface{}) template.HTML {
	var buf bytes.Buffer
	templates.ExecuteTemplate(&buf, page+".html", data)
	return template.HTML(buf.Bytes())
}

//TODO return err
func renderJSON(page string) template.HTML {
	contents, _ := ioutil.ReadFile(page)

	var data interface{}
	json.Unmarshal(contents, &data)

	page = filepath.Base(page)
	page = page[:len(page)-len(filepath.Ext(page))]

	return renderHTML(page, data)
}

//TODO return err
func renderMarkdown(page string) template.HTML {
	md, _ := ioutil.ReadFile(page)
	return template.HTML(blackfriday.MarkdownCommon(md))
}

//reload templates on modify
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
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", handleContent())
	http.HandleFunc("/.well-known/acme-challenge/", acmeChallenge)
	serveFile("/favicon.ico", "./favicon.ico")
	serveFile("/sitemap.xml", "./sitemap.xml")
	serveFile("/robots.txt", "./robots.txt")
	// localhost:5000
	go listenForChanges()
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func acmeChallenge(w http.ResponseWriter, r *http.Request) {
	pt := strings.TrimPrefix(r.URL.Path, "/.well-known/acme-challenge/")
	rk := ""

	k := os.Getenv("ACME_KEY")
	t := os.Getenv("ACME_TOKEN")
	if k != "" && t != "" {
		if pt == t {
			rk = k
		}
	} else {
		for i := 1; ; i++ {
			is := strconv.Itoa(i)
			k = os.Getenv("ACME_KEY_" + is)
			t = os.Getenv("ACME_TOKEN_" + is)
			if k != "" && t != "" {
				if pt == t {
					rk = k
					break
				}
			} else {
				break
			}
		}
	}

	if rk != "" {
		fmt.Fprint(w, rk)
	} else {
		http.NotFound(w, r)
	}
}
