package main

import (
	"fmt"
	"golangWebApp/pkg/config"
	"golangWebApp/pkg/handlers"
	"golangWebApp/pkg/render"
	"log"
	"net/http"
)

const portNumber string = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache:", err)

	}
	app.TemplateCache = tc
	app.UseCache = true

	//send app to handlers & get it back to send it to each individual func
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//Routes
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
