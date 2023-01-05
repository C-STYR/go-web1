package main

import (
	"testing"

	"github.com/C-STYR/go-web1/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	// create a multiplexer for testing
	mux := routes(&app)

	switch v := mux.(type) {

	case *chi.Mux:
	default:
		t.Errorf("Type is not not chi.Mux but %T", v)
	}
}
