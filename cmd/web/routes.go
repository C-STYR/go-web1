package main

import (
	"net/http"

	"github.com/C-STYR/go-web1/pkg/config"
	"github.com/C-STYR/go-web1/pkg/handlers"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
