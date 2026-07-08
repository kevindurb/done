package app

import (
	"database/sql"
	"log"

	"github.com/gorilla/sessions"
	"github.com/kevindurb/done/config"
	"github.com/kevindurb/done/migrations"
	"github.com/kevindurb/done/sqlcgen"
)

type App struct {
	s  sessions.Store
	db *sql.DB
	q  *sqlcgen.Queries
}

func New(c *config.Config) *App {
	db, err := sql.Open("sqlite", c.DBPath)

	if err != nil {
		log.Panicf("Error opening db: %v", err)
	}

	err = migrations.Up(db)
	if err != nil {
		log.Panicf("Error migrating db: %v", err)
	}

	q := sqlcgen.New(db)

	s := sessions.NewCookieStore([]byte(c.SecretKey))
	return &App{s, db, q}
}
