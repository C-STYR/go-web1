package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/C-STYR/go-web1/internal/config"
	"github.com/C-STYR/go-web1/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	// to store non-primitives in the session
	gob.Register(models.Reservation{})

	// change to true in production
	testApp.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	// set session parameters
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// assign the session to this config to expose it to middleware
	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

// creating an interface to satisfy the requirements of http.ResponseWriter
type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(statusCode int) {
	// no need to return anything
}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
