package main

import (
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

var templates = template.Must(template.ParseFiles("tmpl/footer.html", "tmpl/header.html", "tmpl/nep_recruit.html", "tmpl/episode_view.html", "tmpl/map.html", "tmpl/menubar.html"));

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", "some_string")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// epview regex used to be "/watch/?:show/:season([0-9]+)/:episode([0-9]+)"
var validPath = regexp.MustCompile("/|^/watch/divisionrune/([0-9]+)/([0-9]+)$")
//var validPath = regexp.MustCompile("^/(episode_view|watch)/([a-zA-Z0-9]+)$")

func episodeHandler(w http.ResponseWriter, r *http.Request, page string) {
	renderTemplate(w, "episode_view")
}

func mainHandler(w http.ResponseWriter, r *http.Request, page string) {
	renderTemplate(w, "nep_recruit")
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
    some_string := strings.Split(r.URL.Path[1:], "/")
    //some_string := r.URL.Path[1:]
		fn(w, r, some_string[1])
	}
}

func main() {
	http.HandleFunc("/", makeHandler(mainHandler))
	http.HandleFunc("/watch/", makeHandler(episodeHandler))

	http.ListenAndServe(":8080", nil)
}
