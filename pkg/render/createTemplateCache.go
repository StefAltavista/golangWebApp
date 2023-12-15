package render

import (
	"log"
	"path/filepath"
	"text/template"
)

func CreateTemplateCache() (map[string]*template.Template, error) {
	log.Println("creating template cache")

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.ParseFiles(page, "./templates/base.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		myCache[name] = t
	}

	return myCache, nil
}
