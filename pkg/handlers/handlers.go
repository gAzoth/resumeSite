package handlers

import (
	"net/http"

	"github.com/gAzoth/resumeSite/pkg/config"
	"github.com/gAzoth/resumeSite/pkg/models"
	"github.com/gAzoth/resumeSite/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository if the new repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform logic
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
