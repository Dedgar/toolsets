package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

func getShow(c echo.Context) error {
	show := c.Param("show")
	season := c.Param("season")
	episode := c.Param("episode")
	//	return c.String(http.StatusOK, show+season+episode)
	//	c.SetParamNames("show")
	//	c.SetParamValues(c.QueryParam("show"))
	return c.Render(http.StatusOK, "episode_view.html", map[string]interface{}{
		"show":    show,
		"season":  season,
		"episode": episode,
	})
}

func getMain(c echo.Context) error {
	return c.Render(http.StatusOK, "nep_recruit.html", "main")
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("tmpl/*.html")),
	}
	e := echo.New()
	e.Static("/", "static")
	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", getMain)
	//e.GET("/watch/:show", getShow)
	e.GET("/watch/:show/:season/:episode", getShow)
	e.Logger.Info(e.Start(":1323"))
}
