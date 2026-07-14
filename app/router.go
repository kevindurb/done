package app

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/login", func(r chi.Router) {
		r.Get("/", a.loginShow)
		r.Post("/", a.login)
	})

	r.Route("/signup", func(r chi.Router) {
		r.Get("/", a.signupShow)
		r.Post("/", a.signup)
	})

	r.With(a.requireAuth).Get("/", a.home)

	r.With(a.requireAuth).Route("/tasks", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", a.tasksShow)
			r.Post("/done", a.tasksMarkDone)
			r.Post("/delete", a.tasksDelete)
		})

		r.Get("/", a.tasksList)
		r.Post("/", a.tasksCreate)
		r.Get("/new", a.tasksNew)
		r.Get("/done", a.tasksListDone)
	})

	r.With(a.requireAuth).Route("/projects", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", a.projectsShow)
			r.Post("/delete", a.projectsDelete)
		})

		r.Get("/", a.projectsList)
		r.Post("/", a.projectsCreate)
		r.Get("/new", a.projectsNew)
	})

	return r
}
