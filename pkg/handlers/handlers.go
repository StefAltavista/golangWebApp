package handlers

import (
	"golangWebApp/pkg/config"
	"golangWebApp/pkg/models"
	"golangWebApp/pkg/render"
	"net/http"
)

// holds data sent from handlers to templates

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// routes

func (receiver *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	receiver.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (receiver *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIP := receiver.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["test"] = "test dada"
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
