package pages

import (
	"net/http"

	"github.com/kevindurb/done/html/layouts"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Signup() g.Node {
	return layouts.Layout(
		h.H1(g.Text("Signup")),
		h.Form(
			h.Method(http.MethodPost),
			h.Action("/signup"),
			h.Div(
				h.Label(h.For("email"), g.Text("Email")),
				h.Input(h.Type("email"), h.ID("email"), h.Name("Email")),
			),
			h.Div(
				h.Label(h.For("password"), g.Text("Password")),
				h.Input(h.Type("password"), h.ID("password"), h.Name("Password")),
			),
			h.Button(h.Type("submit"), g.Text("Signup")),
			h.A(h.Href("/login"), g.Text("Login")),
		),
	)
}
