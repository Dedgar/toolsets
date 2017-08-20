package main

import (
	"html/template"
	"net/http"
	"regexp"
	"fmt"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "edit")
}

var templates = template.Must(template.ParseFiles("tmpl/nep_recruit.html"))

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", "nep_recruit")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|watch|view)/([a-zA-Z0-9]+)$")

func regHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "nep_recruit")
}

func staticHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/", staticHandler(mainHandler))
	http.HandleFunc("/edit/", regHandler(editHandler))

  fmt.Println("starting server")
	http.ListenAndServe(":8080", nil)
}
