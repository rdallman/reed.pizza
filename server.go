package main

import (
	"bytes"
	"context"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/howeyc/fsnotify"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/crypto/acme/autocert"
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

// certCache just stores certs in memory, since we prob won't restart this weekly?
type certCache struct {
	sync.RWMutex

	m map[string][]byte
}

// Get returns a certificate data for the specified key.
// If there's no such key, Get returns ErrCacheMiss.
func (c *certCache) Get(ctx context.Context, key string) ([]byte, error) {
	c.RLock()
	v, ok := c.m[key]
	c.RUnlock()
	if !ok {
		return nil, autocert.ErrCacheMiss
	}
	return v, nil
}

// Put stores the data in the cache under the specified key.
// Underlying implementations may use any data storage format,
// as long as the reverse operation, Get, results in the original data.
func (c *certCache) Put(ctx context.Context, key string, data []byte) error {
	c.Lock()
	c.m[key] = data
	c.Unlock()
	return nil
}

// Delete removes a certificate data from the cache under the specified key.
// If there's no such key in the cache, Delete returns nil.
func (c *certCache) Delete(ctx context.Context, key string) error {
	delete(c.m, key)
	return nil
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
		err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("ListenAndServe:", err)
		}
		return
	}

	log.Println("running in prod mode")

	// static functions, for memory/speed
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/hire/resume", handleResume)

	m := &autocert.Manager{
		Cache:      new(certCache),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("rdallman.com", "www.rdallman.com"),
	}
	s := &http.Server{
		Addr:      ":https",
		TLSConfig: m.TLSConfig(),
	}
	err := s.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err)
	}
}
