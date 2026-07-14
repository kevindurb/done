package pages

import (
	"net/http"

	"github.com/kevindurb/done/html/layouts"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Login() g.Node {
	return layouts.Layout(
		h.H1(g.Text("Login")),
		h.Form(
			h.Method(http.MethodPost),
			h.Action("/login"),
			h.Div(
				h.Label(h.For("email"), g.Text("Email")),
				h.Input(h.Type("email"), h.ID("email"), h.Name("Email")),
			),
			h.Div(
				h.Label(h.For("password"), g.Text("Password")),
				h.Input(h.Type("password"), h.ID("password"), h.Name("Password")),
			),
			h.Button(h.Type("submit"), g.Text("Login")),
			h.A(h.Href("/signup"), g.Text("Signup")),
		),
	)
}
