package handlers

import (
	"net/http"

	"github.com/iyke0z/Bookings/pkg/config"
	"github.com/iyke0z/Bookings/pkg/models"
	"github.com/iyke0z/Bookings/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	
}

//NewRepo creates a new reporisitory
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

//NewHandler sets the reporsitory for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	remote_ip := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remote_ip)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
 
// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})


}

