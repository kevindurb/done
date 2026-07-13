package app

import (
	"net/http"
)

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/tasks", http.StatusFound)
}
