package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func manHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, eat %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/food", manHandler)
    http.ListenAndServe(":8080", nil)
}
