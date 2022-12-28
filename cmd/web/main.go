package main

import (
	"fmt"
	"net/http"

	"github.com/C-STYR/go-web1/pkg/config"
	"github.com/C-STYR/go-web1/pkg/handlers"
	"github.com/C-STYR/go-web1/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("MAIN cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
