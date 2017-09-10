package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	//	"strconv"
	//	"strings"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "shino"
	password = "agiridesu"
	dbname   = "kanji"
)

//type Kanji struct {
//	kanj   string
//	von    string
//	vkun   string
//	transl string
//	roma   string
//	rememb string
//	jlpt   string
//	school string
//}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getMain(c echo.Context) error {
	return c.Render(http.StatusOK, "nep_recruit.html", "main")
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

func getLevel(c echo.Context) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	sqlQuery := "SELECT kanj, von, vkun, transl, roma, rememb, jlpt, school FROM info WHERE school = $1"
	rows, err := db.Query(sqlQuery, c.Param("level"))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	//entry := make(map[string]map[string]string)
	//var entry []Kanji
	var entry []string

	for rows.Next() {
		var kanj string
		var von string
		var vkun string
		var transl string
		var roma string
		var rememb string
		var jlpt string
		var school string

		if err := rows.Scan(&kanj, &von, &vkun, &transl, &roma, &rememb, &jlpt, &school); err != nil {
			log.Fatal(err)
		}
		//		entry[kanj] = map[string]string{"kanj": kanj}
		//entry[kanj] = map[string]string{"kanj": kanj, "von": von, "vkun": vkun, "transl": transl, "roma": roma, "rememb": rememb, "jlpt": jlpt, "school": school}
		//		entry = append(entry, Kanji{kanj, von, vkun, transl, roma, rememb, jlpt, school})
		entry = append(entry, kanj)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)

	}

	return c.Render(http.StatusOK, "kanji_list.html", entry) // map[string]interface{}{
	//		"entry": entry,
	//	})

}

func getKanji(c echo.Context) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	// ensure :kanji isn't used as an escaped query like "%e9%9b%a8"
	uni_kanj, err := url.QueryUnescape(c.Param("kanji"))

	if err != nil {
		log.Fatal(err)
	}

	sqlQuery := "SELECT kanj, von, vkun, transl, roma, rememb, jlpt, school FROM info WHERE kanj = $1"
	row := db.QueryRow(sqlQuery, uni_kanj)

	if err != nil {
		log.Fatal(err)
	}

	var kanj string
	var von string
	var vkun string
	var transl string
	var roma string
	var rememb string
	var jlpt string
	var school string
	//	var con_kanji string

	switch err := row.Scan(&kanj, &von, &vkun, &transl, &roma, &rememb, &jlpt, &school); err {
	case sql.ErrNoRows:
		// use a 404 here
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(kanj, von, vkun, transl, roma, rememb, jlpt, school)
	default:
		//  panic(err)
		log.Fatal(err)
	}

	entry := map[string]string{
		"kanj":   kanj,
		"von":    von,
		"vkun":   vkun,
		"transl": transl,
		"roma":   roma,
		"rememb": rememb,
		"jlpt":   jlpt,
		"school": school,
	}
	//		entry = append(entry, Kanji{kanj, von, vkun, transl, roma, rememb, jlpt, school})
	//entry = append(entry, kanj)

	return c.Render(http.StatusOK, "flashcard.html", entry) // map[string]interface{}{
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseFiles("tmpl/map.html",
			"tmpl/kanji_list.html",
			"tmpl/flashcard.html",
			"tmpl/header.html",
			"tmpl/episode_view.html",
			"tmpl/nep_recruit.html",
			"tmpl/footer.html",
			"tmpl/menubar.html",
		)),
	}
	e := echo.New()
	e.Static("/", "static")
	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", getMain)
	//e.GET("/watch/:show", getShow)
	e.GET("/watch/:show/:season/:episode", getShow)
	e.GET("/grade/:level", getLevel)
	e.GET("/jlpt/:level", getLevel)
	e.GET("/kanji/:kanji", getKanji)
	e.Logger.Info(e.Start(":1323"))
}
