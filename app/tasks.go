package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kevindurb/done/html/pages"
	"github.com/kevindurb/done/httpx"
	"github.com/kevindurb/done/sqlcgen"
)

type tasksCreateBody struct {
	Description string `validate:"required"`
	Due         string `validate:"regexp=^\\d\\d\\d\\d-\\d\\d-\\d\\d$"`
	ProjectID   int64
}

func (a *App) tasksList(w http.ResponseWriter, r *http.Request) {
	userID := a.mustUserID(r)
	tasks, err := a.q.ListTasks(r.Context(), userID)
	if err != nil {
		log.Printf("Error getting tasks for user (%d): %v", userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pages.TasksList(pages.TasksListData{Tasks: tasks}).Render(w)
}

func (a *App) tasksListDone(w http.ResponseWriter, r *http.Request) {
	userID := a.mustUserID(r)
	tasks, err := a.q.ListTasksDone(r.Context(), userID)
	if err != nil {
		log.Printf("Error getting done tasks for user (%d): %v", userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pages.TasksList(pages.TasksListData{Tasks: tasks, Done: true}).Render(w)
}

func (a *App) tasksNew(w http.ResponseWriter, r *http.Request) {
	userID := a.mustUserID(r)
	projects, err := a.q.ListProjects(r.Context(), userID)
	if err != nil {
		log.Printf("Error getting projects for user (%d): %v", userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pages.TasksNew(pages.TasksNewData{Projects: projects}).Render(w)
}

func (a *App) tasksCreate(w http.ResponseWriter, r *http.Request) {
	userID := a.mustUserID(r)
	var body tasksCreateBody
	if err := a.fp.Parse(&body, r); err != nil {
		log.Printf("Error parsing body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := a.q.CreateTask(r.Context(), sqlcgen.CreateTaskParams{
		UserID:      userID,
		Description: body.Description,
		ProjectID:   sql.NullInt64{Int64: body.ProjectID, Valid: body.ProjectID > 0},
	})

	if err != nil {
		log.Printf("Error creating task: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/tasks", http.StatusFound)
}

func (a *App) tasksShow(w http.ResponseWriter, r *http.Request) {
	id := httpx.PathInt("id", r)
	userID := a.mustUserID(r)
	task, err := a.q.GetTask(r.Context(), sqlcgen.GetTaskParams{
		ID:     id,
		UserID: userID,
	})

	if err != nil {
		log.Printf("Error getting task (%d) for user (%d): %v", id, userID, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	pages.TasksShow(pages.TasksShowData{Task: task}).Render(w)
}

func (a *App) tasksMarkDone(w http.ResponseWriter, r *http.Request) {
	id := httpx.PathInt("id", r)
	userID := a.mustUserID(r)
	task, err := a.q.GetTask(r.Context(), sqlcgen.GetTaskParams{
		ID:     id,
		UserID: userID,
	})

	if err != nil {
		log.Printf("Error getting task (%d) for user (%d): %v", id, userID, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, err = a.q.MarkTaskDone(r.Context(), sqlcgen.MarkTaskDoneParams{
		ID:     task.ID,
		UserID: userID,
	})

	if err != nil {
		log.Printf("Error marking task (%d) done for user (%d): %v", id, userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/tasks", http.StatusFound)
}
