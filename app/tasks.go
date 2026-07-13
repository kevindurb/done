package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kevindurb/done/html"
	"github.com/kevindurb/done/httpx"
	"github.com/kevindurb/done/sqlcgen"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type tasksCreateBody struct {
	Description string `validate:"required"`
}

func (a *App) tasksList(w http.ResponseWriter, r *http.Request) {
	userID := a.mustUserID(r)
	tasks, err := a.q.ListTasks(r.Context(), userID)
	if err != nil {
		log.Printf("Error getting tasks for user (%d): %v", userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	html.Layout(
		h.H1(g.Text("Tasks")),
		h.A(h.Href("/tasks/new"), g.Text("Add Task")),
		h.Ul(
			g.Map(tasks, func(t sqlcgen.Task) g.Node {
				return h.Li(
					h.Form(
						h.Method(http.MethodPost),
						h.Action(fmt.Sprintf("/tasks/%d/done", t.ID)),
						h.Button(h.Type("submit"), g.Text("Mark Done")),
						h.A(h.Href(fmt.Sprintf("/tasks/%d", t.ID)), g.Text(t.Description)),
					),
				)
			}),
		),
		h.A(h.Href("/tasks/done"), g.Text("Show Done")),
	).Render(w)
}

func (a *App) tasksListDone(w http.ResponseWriter, r *http.Request) {
	userID := a.mustUserID(r)
	tasks, err := a.q.ListTasksDone(r.Context(), userID)
	if err != nil {
		log.Printf("Error getting done tasks for user (%d): %v", userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	html.Layout(
		h.A(h.Href("/tasks"), g.Text("< Back")),
		h.H1(g.Text("Done Tasks")),
		h.Ul(
			g.Map(tasks, func(t sqlcgen.Task) g.Node {
				return h.Li(
					h.A(h.Href(fmt.Sprintf("/tasks/%d", t.ID)), g.Text(t.Description)),
				)
			}),
		),
	).Render(w)
}

func (a *App) tasksNew(w http.ResponseWriter, r *http.Request) {
	html.Layout(
		h.H1(g.Text("New Task")),
		h.Form(
			h.Method(http.MethodPost),
			h.Action("/tasks"),
			h.Input(h.Type("text"), h.Name("Description")),
			h.Button(h.Type("submit"), g.Text("Add")),
		),
	).Render(w)
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

	html.Layout(
		h.A(h.Href("/tasks"), g.Text("< Back")),
		h.H1(g.Text(task.Description)),
	).Render(w)
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
