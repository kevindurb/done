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

	r.With(a.requireAuth).Get("/", a.home)
	r.Get("/login", a.loginShow)
	r.Post("/login", a.login)
	r.Get("/signup", a.signupShow)
	r.Post("/signup", a.signup)

	r.With(a.requireAuth).Route("/tasks", func(r chi.Router) {
		r.Get("/{id}", a.tasksShow)
		r.Post("/{id}/done", a.tasksMarkDone)
		r.Get("/", a.tasksList)
		r.Get("/done", a.tasksListDone)
		r.Get("/new", a.tasksNew)
		r.Post("/", a.tasksCreate)
	})

	r.With(a.requireAuth).Route("/projects", func(r chi.Router) {
		r.Get("/{id}", a.projectsShow)
		r.Get("/", a.projectsList)
		r.Get("/new", a.projectsNew)
		r.Get("/delete", a.projectsDelete)
		r.Post("/", a.projectsCreate)
	})

	return r
}
