package handlers

import (
	"fmt"
	"net/http"

	"github.com/tiangjemuran/bookingapps/pkg/config"
	"github.com/tiangjemuran/bookingapps/pkg/models"
	"github.com/tiangjemuran/bookingapps/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler create ne Handler
func NewHandler(r *Repository) {
	Repo = r
}

//Home handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println("remote ip:", remoteIP)
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Templates(w, "home.page.html", &models.TemplateData{})
}

//About handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "#callBack"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Templates(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
