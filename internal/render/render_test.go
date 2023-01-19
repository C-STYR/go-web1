package render

import (
	"net/http"
	"testing"

	"github.com/C-STYR/go-web1/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	// add some values to the flash variable
	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	// our custom empty interface that implements http.ResponseWriter
	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.html", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}

	// blank.page.html does not exist
	err = RenderTemplate(&ww, r, "blank.page.html", &models.TemplateData{})
	if err == nil {
		t.Error("rendered non-existent template")
	}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}

	// grab the req context
	ctx := r.Context()
	// load the session
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	// add the session data to the original req
	r = r.WithContext(ctx)

	return r, nil

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
