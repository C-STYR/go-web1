package dbrepo

import (
	"database/sql"

	"github.com/C-STYR/go-web1/internal/config"
	"github.com/C-STYR/go-web1/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// the means by which we pass our connection pool and app and return a repository
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}
