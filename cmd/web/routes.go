package main

import (
	"net/http"

	"github.com/iyke0z/Bookings/pkg/config"
	"github.com/iyke0z/Bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler{
	//chiRouterPackage Initialize
	mux := chi.NewRouter()
	
	//A bunch of imported middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	//CSRF Middle ware created by me
	mux.Use(NoSurve)
	//CSRF Middle ware created by me to load session
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//start of loading static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	//end of loading static files
	
	return mux
}

