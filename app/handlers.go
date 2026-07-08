package app

import (
	"net/http"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func (a *App) home(w http.ResponseWriter, r *http.Request) (g.Node, error) {
	return h.HTML(), nil
}
