package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

var baseTemplatePattern = "./templates/base/*.html"
var baseTemplate *template.Template

// Template below is struct implements Renderer interface
type Template struct {
	templates map[string]*template.Template
}

//set render to *Echo
func SetRender(e *echo.Echo) {
	template := &Template{make(map[string]*template.Template)}
	template.registerTemplates()
	e.Renderer = template
}

// implementation of Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].Execute(w, data)
}

// register template in Template.templates
func (t *Template) registerTemplates() {
	//set base template
	baseTemplate = template.Must(template.ParseGlob(baseTemplatePattern))

	t.templates["top"] = parseGlobWithBase("./templates/top/*.html")
	t.templates["products"] = parseGlobWithBase("./templates/products/*.html")

}

//
func parseGlobWithBase(pattern string) *template.Template {
	return template.Must(template.Must(baseTemplate.Clone()).ParseGlob(pattern))
}
