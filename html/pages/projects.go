package pages

import (
	"fmt"
	"net/http"

	"github.com/kevindurb/done/html/components"
	"github.com/kevindurb/done/html/layouts"
	"github.com/kevindurb/done/sqlcgen"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type ProjectsListData struct {
	Projects []sqlcgen.Project
}

func ProjectsList(data ProjectsListData) g.Node {
	return layouts.Layout(
		h.A(h.Href("/"), g.Text("< Back")),
		h.H1(g.Text("Projects")),
		h.A(h.Href("/projects/new"), g.Text("Add Project")),
		h.Ul(
			g.Map(data.Projects, func(t sqlcgen.Project) g.Node {
				return h.Li(
					h.A(h.Href(fmt.Sprintf("/projects/%d", t.ID)), g.Text(t.Name)),
				)
			}),
		),
	)
}

func ProjectsNew() g.Node {
	return layouts.Layout(
		h.A(h.Href("/projects"), g.Text("< Back")),
		h.H1(g.Text("New Project")),
		h.Form(
			h.Method(http.MethodPost),
			h.Action("/projects"),
			h.Input(h.Type("text"), h.Name("Name")),
			h.Button(h.Type("submit"), g.Text("Add")),
		),
	)
}

type ProjectsShowData struct {
	Project sqlcgen.Project
	Tasks   []sqlcgen.Task
}

func ProjectsShow(data ProjectsShowData) g.Node {
	return layouts.Layout(
		h.A(h.Href("/projects"), g.Text("< Back")),
		h.H1(g.Text(data.Project.Name)),
		components.TasksList(components.TasksListData{Tasks: data.Tasks}),
	)
}
