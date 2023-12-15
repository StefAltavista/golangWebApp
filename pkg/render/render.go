package render

import (
	"bytes"
	"golangWebApp/pkg/config"
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

func RenderTemplate(w http.ResponseWriter, tmpl string) {
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
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println("error executing template, ", err)
	}

}
