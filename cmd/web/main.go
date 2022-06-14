package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iyke0z/Bookings/pkg/config"
	"github.com/iyke0z/Bookings/pkg/handlers"
	"github.com/iyke0z/Bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

//Main is the main application function
func main() { 	
	
	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime =24 * time.Hour //the timeline of the session
	session.Cookie.Persist = false //if you want the cookie to dissappear on browser load
	session.Cookie.SameSite = http.SameSiteLaxMode 
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create templace cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app) 
	
	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}