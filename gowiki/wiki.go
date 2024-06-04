package main

//https://go.dev/doc/articles/wiki/

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

const TMPL_PATH = "tmpl/"
const DATA_PATH = "data/"
const HOME_PATH = "/view/FrontPage"

var templates = template.Must(template.ParseFiles(TMPL_PATH+"edit.html", TMPL_PATH+"view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view|static)/([a-zA-Z0-9]+)$")
var staticPath = regexp.MustCompile("^/(static)/([a-zA-Z0-9]+)/([a-zA-Z0-9]+)\\.([a-zA-Z]+)$")

func main() {
	// http.HandleFunc("/static/", makeHandler(staticResHandler))
	http.Handle("/static/", http.FileServer(http.Dir("")))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ********* handler start **************/

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if nil != err {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, HOME_PATH, http.StatusFound)
}

// ********* handler end **************/
func (p *Page) save() error {
	filename := DATA_PATH + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := DATA_PATH + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fmt.Println(m)
		fn(w, r, m[2])
	}
}
