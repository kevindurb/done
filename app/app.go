package app

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"github.com/kevindurb/done/config"
	"github.com/kevindurb/done/migrations"
	"github.com/kevindurb/done/sqlcgen"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	ghttp "maragu.dev/gomponents/http"
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

func (a *App) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", ghttp.Adapt(a.home))

	return r
}
