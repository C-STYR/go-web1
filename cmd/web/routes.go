package main

import (
	"net/http"

	"github.com/C-STYR/go-web1/pkg/config"
	"github.com/C-STYR/go-web1/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// out of the box MW
	mux.Use(middleware.Recoverer)
	
	// Adds CSRF protection to post requests
	mux.Use(NoSurf)

	// Loads and saves the session on every request
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	
	return mux
}
