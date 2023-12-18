package main

import (
	"fmt"
	"golangWebApp/pkg/config"
	"golangWebApp/pkg/handlers"
	"golangWebApp/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change to true when in prod mode
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true //after closing page
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //http vs https

	app.Session = session

	//create template cache
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
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting application on port: ", portNumber)

}
