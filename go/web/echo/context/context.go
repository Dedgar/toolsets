package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	c.Set("key", "val")
	return c.Render(http.StatusOK, "hello", c)
}

func main() {
	e := echo.New()
	tmpl, _ := template.New("test").Parse(`{{define "hello"}}Hello, {{ .Get "key"}}!{{end}}`)
	t := &Template{
		templates: tmpl,
	}
	e.SetRenderer(t)
	e.GET("/hello", Hello)
	log.Fatal(e.Run(standard.New(":1323")))
}
