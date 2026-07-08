package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kevindurb/done/html"
	"github.com/kevindurb/done/sqlcgen"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type tasksCreateBody struct {
	Description string `validate:"required"`
}

func (a *App) tasksList(w http.ResponseWriter, r *http.Request) (g.Node, error) {
	userID := a.mustUserID(r)
	tasks, err := a.q.ListTasks(r.Context(), userID)
	if err != nil {
		log.Printf("Error getting tasks for user (%d): %v", userID, err)
	}

	return html.Layout(
		h.H1(g.Text("Tasks")),
		h.A(h.Href("/tasks/new"), g.Text("Add Task")),
		h.Ul(
			g.Map(tasks, func(t sqlcgen.Task) g.Node {
				return h.Li(g.Text(t.Description))
			}),
		),
	), nil
}

func (a *App) tasksNew(w http.ResponseWriter, r *http.Request) (g.Node, error) {
	return html.Layout(
		h.H1(g.Text("New Task")),
		h.Form(
			h.Method(http.MethodPost),
			h.Action("/tasks"),
			h.Input(h.Type("text"), h.Name("Description")),
			h.Button(h.Type("submit"), g.Text("Add")),
		),
	), nil
}

func (a *App) tasksCreate(w http.ResponseWriter, r *http.Request) {
	userID := a.mustUserID(r)
	var body tasksCreateBody
	if err := a.fp.Parse(&body, r); err != nil {
		log.Printf("Error parsing body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task, err := a.q.CreateTask(r.Context(), sqlcgen.CreateTaskParams{
		UserID:      userID,
		Description: body.Description,
	})

	if err != nil {
		log.Printf("Error creating task: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/tasks/%d", task.ID), http.StatusFound)
}
