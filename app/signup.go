package app

import (
	"net/http"

	"github.com/kevindurb/done/html"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func (a *App) signupShow(w http.ResponseWriter, r *http.Request) (g.Node, error) {
	return html.Layout(
		h.H1(g.Text("Login")),
	), nil
}

func (a *App) signup(w http.ResponseWriter, r *http.Request) {
}
