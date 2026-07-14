package pages

import (
	"net/http"
	"strconv"

	c "github.com/kevindurb/done/html/components"
	"github.com/kevindurb/done/html/layouts"
	"github.com/kevindurb/done/sqlcgen"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type TasksListData struct {
	Tasks []sqlcgen.Task
	Done  bool
}

func TasksList(data TasksListData) g.Node {
	backPath := "/"
	if data.Done {
		backPath = "/tasks"
	}

	return layouts.Layout(
		h.A(h.Href(backPath), g.Text("< Back")),
		h.H1(g.Text("Tasks")),
		h.A(h.Href("/tasks/new"), g.Text("Add Task")),
		c.TasksList(c.TasksListData{Tasks: data.Tasks, Done: data.Done}),
		g.If(!data.Done,
			h.A(h.Href("/tasks/done"), g.Text("Show Done")),
		),
	)
}

type TasksNewData struct {
	Projects []sqlcgen.Project
}

func TasksNew(data TasksNewData) g.Node {
	return layouts.Layout(
		h.A(h.Href("/tasks"), g.Text("< Back")),
		h.H1(g.Text("New Task")),
		h.Form(
			h.Method(http.MethodPost),
			h.Action("/tasks"),
			h.Div(
				h.Label(h.For("ProjectID"), g.Text("Project")),
				h.Select(
					h.Name("ProjectID"),
					h.ID("ProjectID"),
					h.Option(h.Value("0"), g.Text("Default")),
					g.Map(data.Projects, func(p sqlcgen.Project) g.Node {
						return h.Option(h.Value(strconv.FormatInt(p.ID, 10)), g.Text(p.Name))
					}),
				),
			),
			h.Div(
				h.Label(h.For("Description"), g.Text("Description")),
				h.Input(h.Type("text"), h.Required(), h.ID("Description"), h.Name("Description")),
			),
			h.Div(
				h.Label(h.For("Due"), g.Text("Due")),
				h.Input(h.Type("date"), h.ID("Due"), h.Name("Due")),
			),
			h.Div(
				h.Button(h.Type("submit"), g.Text("Add")),
			),
		),
	)
}

type TasksShowData struct {
	Task sqlcgen.Task
}

func TasksShow(data TasksShowData) g.Node {
	return layouts.Layout(
		h.A(h.Href("/tasks"), g.Text("< Back")),
		h.H1(g.Text(data.Task.Description)),
	)
}
