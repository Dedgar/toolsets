package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func getShow(c echo.Context) error {
	episode := c.Param("episode")

	return c.String(http.StatusOK, episode)
}

func main() {
	e := echo.New()
	e.GET("/watch/:show/:season/:episode", getShow)
	e.Logger.Fatal(e.Start(":1323"))
}
