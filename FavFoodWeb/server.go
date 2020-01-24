package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

var baseTemplatePattern = "public/template/base/*.html"

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func ParseGlobWithBase(pattern string) *template.Template {
	baseTemplate := template.Must(template.ParseGlob(baseTemplatePattern))
	desiredTemplate := template.Must(baseTemplate.ParseGlob(pattern))
	return desiredTemplate
}

func main() {
	temp := &Template{
		templates: template.Must(template.ParseGlob("public/template/*.html")),
	}

	e := echo.New()
	e.Renderer = temp

	e.GET("/hello", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", "ZENZAI TABE TAAAAI!!!!!")
	})

	e.GET("/hello2", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", "OSHIRUKO DEMO IIYOOOOOOOOO!!!!!!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
