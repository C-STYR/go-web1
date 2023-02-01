package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/C-STYR/go-web1/internal/config"
	"github.com/C-STYR/go-web1/internal/driver"
	"github.com/C-STYR/go-web1/internal/handlers"
	"github.com/C-STYR/go-web1/internal/models"
	"github.com/C-STYR/go-web1/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session = scs.New()
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	// this must be done in main() rather than run() or else db will close immediately
	defer db.SQL.Close()

	fmt.Printf("starting1 application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// to store non-primitives in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	// change to true in production
	app.InProduction = false

	// set session parameters
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// assign the session to this config to expose it to middleware
	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=colestyron password=")
	if err != nil {
		log.Fatal("Cannot connect to database...dead.")
	}
	log.Println("Database connected.")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("MAIN cannot create template cache", err)
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	return db, nil
}
