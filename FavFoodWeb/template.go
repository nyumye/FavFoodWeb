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
	err := t.templates[name].Execute(w, t.templatesData[name])
	return err
}

// register template in Template.templates
func (t *Template) registerTemplates() {
	//set base template
	baseTemplate = template.Must(template.ParseGlob(baseTemplatePattern))
	baseTemplateDatas = makeBaseTemplateDatas()

	t.templates["top"] = parseGlobWithBase("./templates/top/*.html")
	t.templatesData["top"] = new(templateDatas).registerCssJsDatasWithBase("top")

	t.templates["products"] = parseGlobWithBase("./templates/products/*.html")
	t.templatesData["products"] = new(templateDatas).registerCssJsDatasWithBase("products").registerAllFoodDocument()

	t.templates["fortune"] = parseGlobWithBase("./templates/fortune/*.html")
	t.templatesData["fortune"] = new(templateDatas).registerCssJsDatasWithBase("fortune")

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
	Foods []foodDataModel
}

func makeBaseTemplateDatas() *templateDatas {
	return &templateDatas{
		Csses: readPublicDir("base/css", ".css"),
		Jses:  readPublicDir("base/js", ".js"),
		// Images: readPublicDir(folderName + "/image"),
	}
}

func (tempDatas *templateDatas) registerCssJsDatasWithBase(folderName string) *templateDatas {
	tempDatas.Csses = append(baseTemplateDatas.Csses, readPublicDir(folderName+"/css", ".css")...)
	tempDatas.Jses = append(baseTemplateDatas.Jses, readPublicDir(folderName+"/js", ".js")...)
	return tempDatas
	// return &templateDatas{
	// 	Csses: append(baseTemplateDatas.Csses, readPublicDir(folderName+"/css", ".css")...),
	// 	Jses:  append(baseTemplateDatas.Jses, readPublicDir(folderName+"/js", ".js")...),
	// 	// Images: append(baseTemplateDatas.Images, readPublicDir(folderName+"/image")...),
	// }
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

// set all food data in database
func (tempDatas *templateDatas) registerAllFoodDocument() *templateDatas {
	tempDatas.Foods = findAllFoodDocument()
	return tempDatas
}
