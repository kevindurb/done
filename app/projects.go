package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kevindurb/done/html/pages"
	"github.com/kevindurb/done/httpx"
	"github.com/kevindurb/done/sqlcgen"
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

	pages.ProjectsList(pages.ProjectsListData{
		Projects: projects,
	}).Render(w)
}

func (a *App) projectsNew(w http.ResponseWriter, r *http.Request) {
	pages.ProjectsNew().Render(w)
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

	pages.ProjectsShow(pages.ProjectsShowData{
		Project: project,
		Tasks:   tasks,
	}).Render(w)
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
