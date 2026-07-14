package components

import (
	"fmt"
	"net/http"

	"github.com/kevindurb/done/sqlcgen"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type TasksListData struct {
	Tasks []sqlcgen.Task
	Done  bool
}

func TasksList(data TasksListData) g.Node {
	return h.Ul(
		g.Map(data.Tasks, func(t sqlcgen.Task) g.Node {
			return h.Li(
				h.Form(
					h.Method(http.MethodPost),
					h.Action(fmt.Sprintf("/tasks/%d/done", t.ID)),
					g.If(!data.Done,
						h.Button(h.Type("submit"), g.Text("Mark Done")),
					),
					g.If(t.Due.Valid, g.Text(t.Due.String)),
					h.A(h.Href(fmt.Sprintf("/tasks/%d", t.ID)), g.Text(t.Description)),
				),
			)
		}),
	)
}
