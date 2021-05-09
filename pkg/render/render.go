package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tiangjemuran/bookingapps/pkg/config"
	"github.com/tiangjemuran/bookingapps/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

//NewTemplates is assign template
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

//Templates for handler
func Templates(w http.ResponseWriter, page string, td *models.TemplateData) {
	//get te template cache from the app config
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, isExists := tc[page]
	if !isExists {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// CreateTemplateCache is cretae template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("error create template : ", err)
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			fmt.Println("error macthes : ", err)
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				fmt.Println("error macthes parse : ", err)
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
