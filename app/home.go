package app

import (
	"net/http"

	"github.com/kevindurb/done/html/pages"
)

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(w)
}
