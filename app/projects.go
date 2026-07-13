package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/kevindurb/done/html"
	"github.com/kevindurb/done/httpx"
	"github.com/kevindurb/done/sqlcgen"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type projectsCreateBody struct {
	Name string `validate:"required"`
}

func (a *App) projectsList(w http.ResponseWriter, r *http.Request) {
	userID := a.mustUserID(r)
	projects, err := a.q.ListProjects(r.Context(), userID)
	if err != nil {
		log.Printf("Error getting projects for user (%d): %v", userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	html.Layout(
		h.A(h.Href("/"), g.Text("< Back")),
		h.H1(g.Text("Projects")),
		h.A(h.Href("/projects/new"), g.Text("Add Project")),
		h.Ul(
			g.Map(projects, func(t sqlcgen.Project) g.Node {
				return h.Li(
					h.A(h.Href(fmt.Sprintf("/projects/%d", t.ID)), g.Text(t.Name)),
				)
			}),
		),
	).Render(w)
}

func (a *App) projectsNew(w http.ResponseWriter, r *http.Request) {
	html.Layout(
		h.H1(g.Text("New Project")),
		h.Form(
			h.Method(http.MethodPost),
			h.Action("/projects"),
			h.Input(h.Type("text"), h.Name("Name")),
			h.Button(h.Type("submit"), g.Text("Add")),
		),
	).Render(w)
}

func (a *App) projectsCreate(w http.ResponseWriter, r *http.Request) {
	userID := a.mustUserID(r)
	var body projectsCreateBody
	if err := a.fp.Parse(&body, r); err != nil {
		log.Printf("Error parsing body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := a.q.CreateProject(r.Context(), sqlcgen.CreateProjectParams{
		UserID: userID,
		Name:   body.Name,
	})

	if err != nil {
		log.Printf("Error creating project: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/projects", http.StatusFound)
}

func (a *App) projectsShow(w http.ResponseWriter, r *http.Request) {
	id := httpx.PathInt("id", r)
	userID := a.mustUserID(r)
	project, err := a.q.GetProject(r.Context(), sqlcgen.GetProjectParams{
		ID:     id,
		UserID: userID,
	})

	if err != nil {
		log.Printf("Error getting project (%d) for user (%d): %v", id, userID, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tasks, err := a.q.ListTasksByProject(r.Context(), sqlcgen.ListTasksByProjectParams{
		ProjectID: sql.NullInt64{Int64: project.ID, Valid: true},
		UserID:    userID,
	})

	if err != nil {
		log.Printf("Error getting tasks for project (%d) for user (%d): %v", id, userID, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	html.Layout(
		h.A(h.Href("/projects"), g.Text("< Back")),
		h.H1(g.Text(project.Name)),
		html.TasksList(tasks),
	).Render(w)
}

func (a *App) projectsDelete(w http.ResponseWriter, r *http.Request) {
	id := httpx.PathInt("id", r)
	userID := a.mustUserID(r)
	project, err := a.q.GetProject(r.Context(), sqlcgen.GetProjectParams{
		ID:     id,
		UserID: userID,
	})

	if err != nil {
		log.Printf("Error getting project (%d) for user (%d): %v", id, userID, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = a.q.DeleteProject(r.Context(), sqlcgen.DeleteProjectParams{
		ID:     project.ID,
		UserID: userID,
	})

	if err != nil {
		log.Printf("Error deleting project (%d) for user (%d): %v", id, userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/projects", http.StatusFound)
}
