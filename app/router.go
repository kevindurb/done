package app

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	ghttp "maragu.dev/gomponents/http"
)

func (a *App) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", ghttp.Adapt(a.home))
	r.Get("/login", ghttp.Adapt(a.loginShow))
	r.Post("/login", a.login)
	r.Get("/signup", ghttp.Adapt(a.signupShow))
	r.Post("/signup", a.signup)

	return r
}
