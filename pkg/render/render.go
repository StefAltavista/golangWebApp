package render

import (
	"bytes"
	"golangWebApp/pkg/config"
	"golangWebApp/pkg/models"
	"log"
	"net/http"
	"text/template"
)

//Creating template on each reload

// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 		parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 		err := parsedTemplate.Execute(w, nil)

// 		if err != nil {
// 			fmt.Println("error parsing template: ", err)
// 		}
// }

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var TemplateCache map[string]*template.Template
	var err error = nil
	if app.UseCache {
		TemplateCache = app.TemplateCache
	} else {
		TemplateCache, err = CreateTemplateCache()
		if err != nil {
			log.Println("error creating template cache")
		}
	}

	t, exists := TemplateCache[tmpl]

	if !exists {
		log.Fatal("could not find template")

	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	err = t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println("error executing template, ", err)
	}

}
