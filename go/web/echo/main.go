package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// GET /watch/:show/:season/:episode
func getShow(c echo.Context) error {
	show := c.Param("show")
	season := c.Param("season")
	episode := c.Param("episode")

	return c.Render(http.StatusOK, "episode_view.html", map[string]interface{}{
		"show":    show,
		"season":  season,
		"episode": episode,
	})
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
