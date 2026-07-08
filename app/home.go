package app

import (
	"net/http"

	g "maragu.dev/gomponents"
)

func (a *App) home(w http.ResponseWriter, r *http.Request) (g.Node, error) {
	http.Redirect(w, r, "/tasks", http.StatusFound)
	return nil, nil
}
