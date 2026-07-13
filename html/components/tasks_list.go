package components

import (
	"fmt"
	"net/http"

	"github.com/kevindurb/done/sqlcgen"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TasksList(tasks []sqlcgen.Task) g.Node {
	return h.Ul(
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
	)
}
