package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/C-STYR/go-web1/pkg/config"
	"github.com/C-STYR/go-web1/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the render pkg
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache ")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	if err := t.Execute(buf, td); err != nil {
		fmt.Println(err)
	}

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("templates/base.layout.html", page)
		if err != nil {
			fmt.Println(err)
			return myCache, err
		}

		// matches, err := filepath.Glob("./templates/*.layout.html")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return myCache, err
		// }

		// if len(matches) > 0 {
		// 	// fmt.Println("matches found", matches)
		// 	// ParseGlob returns *Templates
		// 	ts, err = template.ParseGlob("./templates/*.layout.html")
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return myCache, err
		// 	}
		// }
		myCache[name] = ts
	}

	return myCache, nil
}
