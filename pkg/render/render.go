package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/iyke0z/Bookings/pkg/config"
	"github.com/iyke0z/Bookings/pkg/models"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

//NewTemplates sets the config for the new template
func NewTemplates(a *config.AppConfig){
	app = a 
} 

func AddDefaultData(td *models.TemplateData) *models.TemplateData{

	return td
}

//RenderTemplate renders template using HTML render
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// get the 
	if app.UseCache { 
		tc = app.TemplateCache  
	}else {
		tc, _ = CreateTemplateCache()
	} 

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

//CreateTemplateCache creates a template cache in a map
func CreateTemplateCache() (map[string]*template.Template, error){
	
	myCache := map[string]*template.Template{}

	//find all pages that ends with .page.tmpl 
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil{
		return myCache, err
	}

	//loop through the pages and add them to the myCache
	for _, page:= range pages {
		name := filepath.Base(page)
		// create a new parsed template files, and you can also attach functions to be used in these template
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err			
		}

		// check if the template matches a layout (base.layout.tmpl)		
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err		
		}
		
		// if seen, the length of matches is greater than 0, then parseGlob the layout
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")		
			if err != nil {
				return myCache, err		
			}
		}

		// assign the mycache[name] to the template set which keeps all the ts in the mycache variable
		myCache[name] = ts
		
	}
	return myCache, nil
}