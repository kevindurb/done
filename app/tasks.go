package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kevindurb/done/html"
	"github.com/kevindurb/done/html/layouts"
	"github.com/kevindurb/done/httpx"
	"github.com/kevindurb/done/sqlcgen"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type tasksCreateBody struct {
	Description string `validate:"required"`
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

	layouts.Layout(
		h.H1(g.Text("Tasks")),
		h.A(h.Href("/tasks/new"), g.Text("Add Task")),
		h.A(h.Href("/projects/new"), g.Text("Add Project")),
		html.TasksList(tasks),
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

	layouts.Layout(
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
	userID := a.mustUserID(r)
	projects, err := a.q.ListProjects(r.Context(), userID)
	if err != nil {
		log.Printf("Error getting projects for user (%d): %v", userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	layouts.Layout(
		h.H1(g.Text("New Task")),
		h.Form(
			h.Method(http.MethodPost),
			h.Action("/tasks"),
			h.Div(
				h.Select(
					h.Name("ProjectID"),
					h.Option(h.Value("0"), g.Text("Default")),
					g.Map(projects, func(p sqlcgen.Project) g.Node {
						return h.Option(h.Value(strconv.FormatInt(p.ID, 10)), g.Text(p.Name))
					}),
				),
			),
			h.Input(h.Type("text"), h.Required(), h.Name("Description")),
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

	layouts.Layout(
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
