package pages

import (
	"github.com/kevindurb/done/html/layouts"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Home() g.Node {
	return layouts.Layout(
		h.H1(g.Text("Home")),
		h.A(h.Href("/tasks"), g.Text("All Tasks")),
		h.A(h.Href("/projects"), g.Text("All Projects")),
	)
}
