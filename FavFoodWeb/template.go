package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
)

var baseTemplatePattern = "./templates/base/*.html"
var baseTemplate *template.Template

// Template below is struct implements Renderer interface
type Template struct {
	templates     map[string]*template.Template
	templatesData map[string]*templateDatas
}

//set render to *Echo
func SetRender(e *echo.Echo) {
	template := &Template{
		templates:     make(map[string]*template.Template),
		templatesData: make(map[string]*templateDatas),
	}
	template.registerTemplates()
	e.Renderer = template
}

// implementation of Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].Execute(w, t.templatesData[name])
}

// register template in Template.templates
func (t *Template) registerTemplates() {
	//set base template
	baseTemplate = template.Must(template.ParseGlob(baseTemplatePattern))
	baseTemplateDatas = makeTemplateDatas("base")

	t.templates["top"] = parseGlobWithBase("./templates/top/*.html")
	t.templatesData["top"] = makeTemplateDatasWithBase("top")

	t.templates["products"] = parseGlobWithBase("./templates/products/*.html")
	t.templatesData["products"] = makeTemplateDatasWithBase("products")

}

//
func parseGlobWithBase(pattern string) *template.Template {
	return template.Must(template.Must(baseTemplate.Clone()).ParseGlob(pattern))
}

//--- template data ---
var baseTemplateDatas *templateDatas

//data objects
type templateDatas struct {
	Csses []string
	Jses  []string
	// Images []string
}

func makeTemplateDatas(folderName string) *templateDatas {
	return &templateDatas{
		Csses: readPublicDir(folderName+"/css", ".css"),
		Jses:  readPublicDir(folderName+"/js", ".js"),
		// Images: readPublicDir(folderName + "/image"),
	}
}

func makeTemplateDatasWithBase(folderName string) *templateDatas {
	return &templateDatas{
		Csses: append(baseTemplateDatas.Csses, readPublicDir(folderName+"/css", ".css")...),
		Jses:  append(baseTemplateDatas.Jses, readPublicDir(folderName+"/js", ".js")...),
		// Images: append(baseTemplateDatas.Images, readPublicDir(folderName+"/image")...),
	}
}

var publicPath = "./public/"

//reads "./public/{{dirname}}" dir
func readPublicDir(dirname, suffix string) []string {
	files, err := ioutil.ReadDir(publicPath + dirname)
	if err != nil {
		log.Fatal(err)
	}

	var pathes []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), suffix) {
			pathes = append(pathes, filepath.Join(".", dirname, file.Name()))
		}
	}
	return pathes
}
