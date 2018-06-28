package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
)

//func visit(path string, f os.FileInfo, err error) error {
//	fmt.Printf("Visited: %s\n", path)
//	return nil
//}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func New() *Template {
	t = &Template{
		templates: func() *template.Template {
			templ := template.New("")
			if err := filepath.Walk("views", func(path string, info os.FileInfo, err error) error {
				if strings.Contains(path, ".html") {
					_, err = templ.ParseFiles(path)
					if err != nil {
						log.Println(err)
					}
				}
				return err
			}); err != nil {
				panic(err)
			}
			return templ
		}(),
	}
}

//func main() {
//	root := flag.Arg(0)

//template := template.New("template")
//t := &Template{templates: template.Must(template.ParseFiles(path))}
//	t := &Template{templates:"test"}

//	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
//		if strings.HasSuffix(path, ".html") {
//			template.Must(template.ParseFiles(path)),
//			template.ParseFiles(path)
//fmt.Printf("Visited: %s\n", path)
//		}

//		return nil
//	}
//	filepath.Walk("./", visit)
//fmt.Printf("filepath.Walk() returned %v\n", err)
//}
