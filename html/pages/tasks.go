package pages

import (
	"net/http"
	"strconv"

	"github.com/kevindurb/done/html/components"
	"github.com/kevindurb/done/html/layouts"
	"github.com/kevindurb/done/sqlcgen"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type TasksListData struct {
	Tasks []sqlcgen.Task
}

func TasksList(data TasksListData) g.Node {
	return layouts.Layout(
		h.A(h.Href("/"), g.Text("< Back")),
		h.H1(g.Text("Tasks")),
		h.A(h.Href("/tasks/new"), g.Text("Add Task")),
		components.TasksList(data.Tasks),
		h.A(h.Href("/tasks/done"), g.Text("Show Done")),
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
				h.Select(
					h.Name("ProjectID"),
					h.Option(h.Value("0"), g.Text("Default")),
					g.Map(data.Projects, func(p sqlcgen.Project) g.Node {
						return h.Option(h.Value(strconv.FormatInt(p.ID, 10)), g.Text(p.Name))
					}),
				),
			),
			h.Input(h.Type("text"), h.Required(), h.Name("Description")),
			h.Button(h.Type("submit"), g.Text("Add")),
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
