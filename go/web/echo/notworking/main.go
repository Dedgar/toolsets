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

type Video interface {
	show() string
	season() string
	episode() string
}

type Show struct {
}

type Season struct {
}

type Episode struct {
}

func (s Show) Video() string {
	return ":show"
}

func (s Season) Video() string {
	return ":season"
}

func (e Episode) Video() string {
	return ":episode"
}

type EP struct {
	show, season, episode string
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getShow(c echo.Context) error {
	show := c.QueryParam("show")
	season := c.QueryParam("season")
	episode := c.QueryParam("episode")
	//  video := EP{c.QueryParam("show"), c.QueryParam("season"), c.QueryParam("episode")}
	//  video := []Video{Show{}, Season{}, Episode{}}
	c.SetParamNames("show")
	c.SetParamValues(c.QueryParam("show"))
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
	//	e.File("/favicon.ico", "/img/favicon.ico")
	e.GET("/", getMain)
	e.GET("/watch/:show/:season/:episode", getShow)
	e.Logger.Info(e.Start(":1323"))
}
