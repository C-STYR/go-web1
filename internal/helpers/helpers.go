package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/C-STYR/go-web1/internal/config"
)

var app *config.AppConfig

// New Helpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of ", status)

	// http.Error replies to the req with status code and error msg
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	// print error and stack trace
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	// write internal server error
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
