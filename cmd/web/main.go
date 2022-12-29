package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/C-STYR/go-web1/pkg/config"
	"github.com/C-STYR/go-web1/pkg/handlers"
	"github.com/C-STYR/go-web1/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig

var session = scs.New()

func main() {

	// change to true in production
	app.InProduction = false

	
	// set session parameters
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	
	// assign the session to this config to expose it to middleware
	app.Session = session
	
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("MAIN cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
