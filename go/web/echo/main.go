package main

import (
	"net/http"
  "html/template"
  "io"
	"github.com/labstack/echo"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func getShow(c echo.Context) error {
//  show := c.QueryParam("show")
//  season := c.QueryParam("season")
//  episode := c.QueryParam("episode")
  return c.Render(http.StatusOK, "test", "episode_view.html")

}

func main() {
  t := &Template{
    templates: template.Must(template.ParseGlob("tmpl/*.html")),
  }
  e := echo.New()
  e.Renderer = t
//	e.GET("/", getShow)
  e.GET("/watch/:show/:season/:episode", getShow)
	e.Logger.Fatal(e.Start(":1323"))
}
